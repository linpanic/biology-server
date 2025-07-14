package api

import (
	"github.com/gin-gonic/gin"
	"github.com/linpanic/biology-server/dto"
	"github.com/linpanic/biology-server/server/service"
)

var (
	UserApi UserAPI
)

type UserAPI struct {
	service.UserService
}

func (u *UserAPI) Register(c *gin.Context) {
	req := new(dto.UserRegisterReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		ReturnJSONError(c, err)
		return
	}
	result := u.UserService.Register(*req)
	ReturnOK(c, result)
	return
}

func (u *UserAPI) Login(c *gin.Context) {
	req := new(dto.UserLoginReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		ReturnJSONError(c, err)
		return
	}
	result := u.UserService.Login(*req)
	ReturnOK(c, result)
	return
}
