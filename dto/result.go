package dto

import (
	"github.com/linpanic/biology-server/cst"
	"net/http"
)

type Result struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Key     string `json:"key,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func NewOKResult(data any) Result {
	return Result{
		Code:    http.StatusOK,
		Message: cst.SUCCESS_MESSAGE,
		Data:    data,
	}
}

func NewErrResult(key, err string) Result {
	if err == "" {
		err = cst.ERROR_MESSAGE
	}
	if key == "" {
		key = cst.UNKNOW_ERROR
	}
	return Result{
		Key:     key,
		Code:    http.StatusBadRequest,
		Message: err,
	}
}

func LoginErrorResult() Result {
	return Result{
		Code:    http.StatusUnauthorized,
		Key:     cst.TOKEN_EXPIRE,
		Message: "登陆身份过期",
	}
}

func TokenErrorResult() Result {
	return Result{
		Code:    -401,
		Key:     cst.NEED_LOGIN,
		Message: "登陆身份过期",
	}
}

func NoTokenResult() Result {
	return Result{
		Code:    -200,
		Key:     cst.NO_TOKEN,
		Message: "no token",
	}
}
