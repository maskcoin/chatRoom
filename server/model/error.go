package model

import "errors"

// 根据业务逻辑，自定义一些错误
var (
	ErrorUserNotExists = errors.New("用户不存在")
	ErrorUserExists    = errors.New("用户已存在")
	ErrorUserPwd       = errors.New("密码不正确")
)
