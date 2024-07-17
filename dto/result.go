package dto

import (
	"github.com/linpanic/biology-server/cst"
	"net/http"
)

type Result struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func NewOKResult(data any) Result {
	return Result{
		Code:    http.StatusOK,
		Message: cst.SUCCESS_MESSAGE,
		Data:    data,
	}
}

func NewErrResult(err string) Result {
	if err == "" {
		err = cst.ERROR_MESSAGE
	}
	return Result{
		Code:    http.StatusBadRequest,
		Message: err,
	}
}

func LoginErrorResult() Result {
	return Result{
		Code:    http.StatusUnauthorized,
		Message: "登陆身份过期",
	}
}
