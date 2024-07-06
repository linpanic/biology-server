package etype

import "errors"

var (
	TimeError        = errors.New("时间错误")
	FileNullError    = errors.New("部分字段不可为空")
	SignError        = errors.New("签名错误")
	NumberEmptyError = errors.New("序列号不存在")
	PageError        = errors.New("分页数据错误")
)
