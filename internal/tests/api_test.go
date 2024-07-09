package tests

import (
	"codearena/internal/db/repository"
	"fmt"
	"testing"
)

func TestInsertOne(t *testing.T) {
	var user = &repository.User{
		Username:   "zql3",
		Email:      "124567890@qq.com",
		Password:   "12345",
		Gender:     "12345",
		ProfilePic: "https://avatar.iran.liara.run/public/boy?username=zql",
	}
	err := repository.InsertOneUser(user)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(user.ID)
}

func TestGetUserList(t *testing.T) {
	userList, err := repository.GetUserList()
	if err != nil {
		t.Error(err.Error())
		return
	}

	for _, user := range userList {
		fmt.Println(user)
	}
}

func TestFindByUsername(t *testing.T) {
	user, err := repository.FindByUsername("zql")
	if err != nil {
		t.Error(err.Error())
		return
	}

	fmt.Println(user.ID)
}
