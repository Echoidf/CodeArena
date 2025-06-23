package repository

import (
	"content-assistant/internal/db/mysql"
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func init() {
	Migrate()
}

type (
	UserStatus byte
	Gender     byte

	User struct {
		ID        uint       `json:"id"`
		Email     string     `json:"email" gorm:"type:varchar(128);not null;comment:邮箱"`
		Username  string     `json:"username" gorm:"type:varchar(32);not null;comment:用户名"`
		Password  string     `json:"password" gorm:"type:varchar(128);not null;comment:密码"`
		Gender    Gender     `json:"gender" gorm:"type:int(2);not null;comment:性别"`
		Avatar    string     `json:"avatar" gorm:"type:varchar(255);not null;comment:头像"`
		Status    UserStatus `json:"actived" gorm:"type:int(2);not null;comment:状态"`
		CreatedAt time.Time  `json:"createdAt"`
		UpdatedAt time.Time  `json:"updatedAt"`
	}
)

const (
	UserStatusInactive UserStatus = iota
	UserStatusActive

	GenderMale Gender = iota
	GenderFemale
)

func InsertOneUser(user *User) error {
	if user == nil {
		return errors.New("user provided must not be nil")
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	res := mysql.DBClient.Create(user)
	if res.Error != nil {
		return res.Error
	}
	zap.L().Info("insert user success", zap.String("username", user.Username))
	return nil
}

func GetUserList() ([]*User, error) {
	var users []*User
	err := mysql.DBClient.Find(&users).Error
	if err != nil {
		zap.L().Error("get user list failed", zap.Error(err))
		return nil, err
	}
	return users, nil
}

func FindByUsername(username string) (*User, error) {
	var user User
	err := mysql.DBClient.Where("username = ?", username).First(&user).Error
	if err != nil {
		zap.L().Error("find user by username failed", zap.Error(err))
		return nil, err
	}
	return &user, nil
}

func FindByUsernameOrEmail(username string, email string) (*User, error) {
	var user User
	err := mysql.DBClient.Where("username = ? OR email = ?", username, email).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		zap.L().Error("find user by username or email failed", zap.String("username", username), zap.String("email", email), zap.Error(err))
		return nil, err
	}
	return &user, nil
}

func FindByUserID(userId uint) (*User, error) {
	var user User
	err := mysql.DBClient.Where("id = ?", userId).First(&user).Error
	if err != nil {
		zap.L().Error("find user by user id failed", zap.Error(err))
		return nil, err
	}
	return &user, nil
}

func DeleteUser(userId uint) error {
	var user User
	if err := mysql.DBClient.Where("id = ?", userId).First(&user).Error; err != nil {
		zap.L().Error("find user by user id failed", zap.Error(err))
		return err
	}

	user.Status = UserStatusInactive
	mysql.DBClient.Save(&user)
	return nil
}

func (user *User) GenerateAvatar() {
	avatar := "https://avatar.iran.liara.run/public"
	if user.Gender == GenderMale {
		avatar += "/boy?username=" + user.Username
	} else {
		avatar += "/girl?username=" + user.Username
	}
	user.Avatar = avatar
}
