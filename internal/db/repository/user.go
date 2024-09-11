package repository

import (
	"codearena/internal/db/mysql"
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id"`
	Phone     string    `json:"phone" gorm:"type:varchar(11);not null;comment:手机号"`
	Username  string    `json:"username" gorm:"type:varchar(32);not null;comment:用户名"`
	Password  string    `json:"password" gorm:"type:varchar(128);not null;comment:密码"`
	Gender    uint8     `json:"gender" gorm:"type:tinyint(1);not null;comment:性别 1-男 2-女"`
	Avatar    string    `json:"avatar" gorm:"type:varchar(255);not null;comment:头像"`
	OpenId    string    `json:"openId" gorm:"type:varchar(128);not null;comment:微信openid"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (User) TableName() string {
	return "sys_user"
}

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
		zap.L().Error("find user by username or email failed", zap.Error(err))
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
	err := mysql.DBClient.Delete(&User{}, "id = ?", userId).Error
	if err != nil {
		zap.L().Error("delete user failed", zap.Error(err))
		return err
	}
	return nil
}
