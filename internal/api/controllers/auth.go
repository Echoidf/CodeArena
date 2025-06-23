package controllers

import (
	"content-assistant/internal/db/repository"
	"content-assistant/pkg/common"
	"content-assistant/pkg/config"
	"content-assistant/pkg/jwt"

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
		common.InternalError(c, err.Error())
		return
	}

	var user *repository.User
	var err error

	if user, err = repository.FindByUsername(req.Username); err != nil {
		common.InternalError(c, err.Error())
		return
	}

	if user == nil {
		common.BadRequest(c, "username or password invalid")
		return
	}

	if user.Status == repository.UserStatusInactive {
		common.BadRequest(c, "account is inactive, please contact admin")
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		common.BadRequest(c, "account or password error")
		return
	}

	generateTokenAndSetCookie(user.ID, c)
	user.Password = ""
	common.Success(c, user)
}

func Signup(c *gin.Context) {
	var req SignupReq
	if err := c.Bind(&req); err != nil {
		common.InternalError(c, err.Error())
		return
	}

	if req.Password != req.ConfirmPassword {
		common.BadRequest(c, "password do not match")
		return
	}

	if req.Email == "" {
		common.BadRequest(c, "email does not provided")
		return
	}

	var user *repository.User
	var err error
	if user, err = repository.FindByUsernameOrEmail(req.Username, req.Email); err != nil {
		common.InternalError(c, err.Error())
		return
	}

	if user.ID != 0 {
		common.Forbidden(c, "account already exists")
		return
	}

	user = &req.User

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	user.GenerateAvatar()
	user.Status = repository.UserStatusActive

	err = repository.InsertOneUser(user)
	if err != nil {
		common.InternalError(c, err.Error())
		return
	}

	common.Success(c, nil)
}

func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", 0, "", "", true, true)
	common.Success(c, "logout successfully")
}

func generateTokenAndSetCookie(userId uint, c *gin.Context) {
	token := jwt.GenerateAccessToken(map[string]any{"id": userId})
	c.SetCookie("auth-token",
		token,
		jwt.AccessTokenExpiredTime,
		"",
		"",
		config.GetConfig().Environment == config.ProductionEnv,
		false)
}
