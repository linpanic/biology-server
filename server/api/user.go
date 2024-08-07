package api

import (
	"github.com/gin-gonic/gin"
	"github.com/linpanic/biology-server/cst"
	"github.com/linpanic/biology-server/dto"
	"github.com/linpanic/biology-server/server/service"
	log "github.com/sirupsen/logrus"
	"net/http"
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
		log.Error(err)
		c.JSON(http.StatusOK, dto.NewErrResult(cst.JSON_ERROR, err.Error()))
		return
	}
	result := u.UserService.Register(*req)
	c.JSON(http.StatusOK, result)
	return
}

func (u *UserAPI) Login(c *gin.Context) {
	req := new(dto.UserLoginReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, dto.NewErrResult(cst.JSON_ERROR, err.Error()))
		return
	}
	result := u.UserService.Login(*req)
	c.JSON(http.StatusOK, result)
	return
}
