package repository

import (
	"codearena/consts"
	"codearena/internal/db/mysql"
	"codearena/pkg/config"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt"
)

type (
	AuthRequest struct {
		Username        string           `json:"username,omitempty"`
		Password        string           `json:"password"`
		Phone           string           `json:"phone,omitempty"`
		MiniProGramCode string           `json:"miniProGramCode,omitempty"`
		LoginType       consts.LoginType `json:"loginType"`
	}

	AuthResponse struct {
		Token    string        `json:"token"`
		UserName string        `json:"username"`
		Avatar   string        `json:"avatar"`
		Gender   consts.Gender `json:"gender"`
	}

	WXLoginResp struct {
		OpenId     string `json:"openid"`      // 用户唯一标识
		SessionKey string `json:"session_key"` // 会话密钥
		UnionId    string `json:"unionid"`     // 用户在微信开放平台账号下的唯一标识
		ErrCode    int    `json:"errcode"`
		ErrMsg     string `json:"errmsg"`
	}
)

var authActionMap map[consts.LoginType]func(authForm *AuthRequest) (user *User, err error)

func init() {
	authActionMap = make(map[consts.LoginType]func(authForm *AuthRequest) (user *User, err error))
	authActionMap[consts.LOGIN_BY_PWD] = LoginByPwd
	authActionMap[consts.LOGIN_BY_WECHAT] = LoginByWechat
}

func Login(authReq *AuthRequest) (res *AuthResponse, err error) {
	var user *User
	fc, ok := authActionMap[authReq.LoginType]
	if ok {
		user, err = fc(authReq)
		if err != nil {
			return nil, err
		}
		return &AuthResponse{
			Token:    GetToken(user),
			UserName: user.Username,
			Avatar:   user.Avatar,
			Gender:   consts.Gender(user.Gender),
		}, nil
	}
	return nil, errors.New(consts.ErrInvalidLoginType)
}

func LoginByPwd(authReq *AuthRequest) (user *User, err error) {
	mysql.DBClient.Table(User{}.TableName()).Where("username = ? || phone = ?", authReq.Username, authReq.Phone).First(&user)
	if user.Password == authReq.Password {
		return user, nil
	}
	return nil, errors.New(consts.ErrInvalidCredentials)
}

func LoginByWechat(authReq *AuthRequest) (user *User, err error) {
	cfg := config.GetConfig().Wechat
	wxResp, err := wxLogin(cfg.AppID, cfg.AppSecret, authReq.MiniProGramCode)
	if err != nil {
		log.Errorf("login by wechat failed, err:%v", err)
		return nil, err
	}
	// query user by openid
	mysql.DBClient.Table(User{}.TableName()).Where("open_id = ?", wxResp.OpenId).First(&user)
	if user.OpenId == "" {
		// create new user
		user = &User{
			OpenId:   wxResp.OpenId,
			Username: "佚名" + wxResp.OpenId[0:6],
			Gender:   uint8(consts.FEMALE),
		}
		mysql.DBClient.Table(User{}.TableName()).Create(user)
	}
	return user, nil
}

func wxLogin(appId, secret, code string) (*WXLoginResp, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	url = fmt.Sprintf(url, appId, secret, code)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	wxResp := &WXLoginResp{}
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&wxResp); err != nil {
		return nil, err
	}

	if wxResp.ErrCode != 0 {
		return nil, fmt.Errorf("errCode:%v  errMsg:%s", wxResp.ErrCode, wxResp.ErrMsg)
	}
	return wxResp, nil
}

func GetToken(user *User) string {
	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.GetConfig().AuthSecret))
	if err != nil {
		log.Error(err)
		return ""
	}

	return t
}
