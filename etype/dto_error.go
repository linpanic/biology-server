package etype

import "errors"

var (
	TimeError     = errors.New("时间错误")
	FileNullError = errors.New("部分字段不可为空")
	SignError     = errors.New("签名错误")
)
