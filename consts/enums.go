package consts

type LoginType uint8

const (
	LOGIN_BY_PWD    LoginType = 1
	LOGIN_BY_WECHAT LoginType = 2
)

type Gender uint8

const (
	MALE   Gender = 1
	FEMALE Gender = 2
)
