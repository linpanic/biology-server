package api

import (
	"github.com/gin-gonic/gin"
	"github.com/linpanic/biology-server/cst"
	"github.com/linpanic/biology-server/dto"
	"github.com/linpanic/biology-server/server/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type AlleleAPI struct {
	service.AlleleService
}

var (
	AlleleApi = new(AlleleAPI)
)

func (a *AlleleAPI) Add(c *gin.Context) {
	req := new(dto.AlleleAddReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, dto.NewErrResult(cst.JSON_ERROR, err.Error()))
		return
	}
	result := a.AlleleService.Add(*req, c.GetInt64("user"))
	c.JSON(http.StatusOK, result)
	return
}

func (a *AlleleAPI) List(c *gin.Context) {
	req := new(dto.AlleleListReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, dto.NewErrResult(cst.JSON_ERROR, err.Error()))
		return
	}
	result := a.AlleleService.AlleleAllSearch(*req)
	c.JSON(http.StatusOK, result)
	return
}

func (a *AlleleAPI) Update(c *gin.Context) {
	req := new(dto.AlleleUpdateReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, dto.NewErrResult(cst.JSON_ERROR, err.Error()))
		return
	}
	result := a.AlleleService.Update(*req, c.GetInt64("user"))
	c.JSON(http.StatusOK, result)
	return
}

func (a *AlleleAPI) Delete(c *gin.Context) {
	req := new(dto.AlleleDelReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, dto.NewErrResult(cst.JSON_ERROR, err.Error()))
		return
	}
	result := a.AlleleService.Delete(*req, c.GetInt64("user"))
	c.JSON(http.StatusOK, result)
	return
}
