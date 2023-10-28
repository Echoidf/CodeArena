package api

import (
	"CodeArena/consts"
	"CodeArena/models"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

func init() {
	authActionMap[consts.LOGIN_BY_MINIPROGRAM] = GetUserByPwd
}

type WXLoginResp struct {
	OpenId     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionId    string `json:"unionid"`     // 用户在微信开放平台账号下的唯一标识
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

func GetUserByMiniProgram(authForm *AuthForm) (user *models.User, err error) {
	resp, err := wxLogin(authForm.AppId, authForm.Secret, authForm.MiniProGramCode)
	if err != nil {
		zap.L().Error(fmt.Sprintf("login by miniapp failed, err:%v", err.Error()))
		return nil, err
	}

	// 查询用户信息
	user, err = models.GetModelByFields[models.User]("user", map[string]interface{}{
		"open_id": resp.OpenId,
	})
	if err != nil {
		zap.L().Error(fmt.Sprintf("query user by openId failed, err:%v", err.Error()))
		return nil, err
	}

	// 用户不存在则创建新用户
	if user == nil {
		user = &models.User{}
		//TODO...
	}
	return
}

func wxLogin(appId, secret, code string) (*WXLoginResp, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	// 合成url, 这里的appId和secret是在微信公众平台上获取的
	url = fmt.Sprintf(url, appId, secret, code)

	// 创建http get请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 解析http请求中body 数据到我们定义的结构体中
	wxResp := &WXLoginResp{}
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&wxResp); err != nil {
		return nil, err
	}

	// 判断微信接口返回的是否是一个异常情况
	if wxResp.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("ErrCode:%v  ErrMsg:%s", wxResp.ErrCode, wxResp.ErrMsg))
	}
	return wxResp, nil
}
