package consts

type LoginType uint8

const (
	LOGIN_BY_PWD         LoginType = 1
	LOGIN_BY_PHONE       LoginType = 2
	LOGIN_BY_MINIPROGRAM LoginType = 3
)
