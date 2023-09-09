package models

import (
	"CodeArena/utils"
	"go.uber.org/zap"
	"time"
)

type User struct {
	Id        int64     `xorm:"bigint(20) notnull pk autoincr" json:"id,omitempty"`
	Username  string    `xorm:"varchar(50) notnull" json:"username,omitempty"`
	Password  string    `xorm:"varchar(255) notnull" json:"password,omitempty"`
	Salt      string    `xorm:"varchar(20) notnull" json:"salt,omitempty"`
	Email     string    `xorm:"varchar(100)" json:"email,omitempty"`
	CreatedAt time.Time `xorm:"datetime" json:"createdAt,omitempty"`
	Phone     string    `xorm:"char(15)" json:"phone,omitempty"`
}

func AddUser(user *User) (affected int64, err error) {
	if user == nil {
		zap.L().Error("user is nil")
		return
	}

	if user.Id == 0 {
		user.Id = utils.NextId().(int64)
	}

	affected, err = Engine.InsertOne(user)
	if err != nil {
		zap.L().Error("insert user failed", zap.Error(err))
	}
	return
}
