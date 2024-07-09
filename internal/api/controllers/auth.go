package controllers

import (
	"codearena/internal/db/repository"
	"codearena/pkg/config"
	"codearena/pkg/jwt"
	"codearena/pkg/response"
	"codearena/pkg/utils"
	"codearena/plugins/email"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type (
	LoginReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	SignupReq struct {
		repository.User
		ConfirmPassword string `json:"confirmPassword"`
	}
)

func Login(c *gin.Context) {
	var req LoginReq
	if err := c.Bind(&req); err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	user, err := repository.FindByUsername(req.Username)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	if user == nil {
		response.Error(c, 400, "username or password valid")
		return
	}

	if user.Actived == 0 {
		response.Error(c, 400, "account not activated, please check your email")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		response.Error(c, 400, "password error")
		return
	}

	generateTokenAndSetCookie(user.ID, c)
	user.Password = ""
	response.Success(c, user)
}

func Signup(c *gin.Context) {
	var req SignupReq
	if err := c.Bind(&req); err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	if req.Password != req.ConfirmPassword {
		response.Error(c, 400, "password do not match")
		return
	}

	if req.Email == "" {
		response.Error(c, 400, "email is required")
		return
	}

	user, err := repository.FindByUsernameOrEmail(req.Username, req.Email)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	if user.ID != 0 {
		response.Error(c, 403, "account already exists")
		return
	}

	user = &req.User

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	profilePic := "https://avatar.iran.liara.run/public"
	if req.Gender == "male" {
		profilePic += "/boy?username=" + user.Username
	} else {
		profilePic += "/girl?username=" + user.Username
	}
	user.ProfilePic = profilePic

	code := utils.GenerateRandomString(8)
	user.ActiveCode = code

	err = repository.InsertOneUser(user)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	// 发送邮件验证码
	err = email.SendEmailWithTmpl(req.Email, code)
	if err != nil {
		response.Error(c, 500, "failed to send email,err: "+err.Error())
		return
	}
	response.Success(c, nil)
}

func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", 0, "", "", true, true)
	response.Success(c, "logout successfully")
}

func ActivateAccount(c *gin.Context) {
	code := c.Query("code")
	username := c.Query("username")
	if code == "" || len(code) != 8 || username == "" {
		response.Error(c, 400, "invalid code or username")
		return
	}

	user, err := repository.ActivateAccount(username, code)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}
	user.Password = ""
	generateTokenAndSetCookie(user.ID, c)

	response.Success(c, user)
}

func generateTokenAndSetCookie(userId uint, c *gin.Context) {
	token := jwt.GenerateAccessToken(map[string]interface{}{"id": userId})
	c.SetCookie("auth-token",
		token,
		jwt.AccessTokenExpiredTime,
		"",
		"",
		config.GetConfig().Environment == config.ProductionEnv,
		false)
}
