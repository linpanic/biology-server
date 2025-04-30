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
	alleleSvc *service.AlleleService
}

func NewAlleleAPI(svc *service.AlleleService) *AlleleAPI {
	return &AlleleAPI{
		alleleSvc: svc,
	}
}

func (a *AlleleAPI) Add(c *gin.Context) {
	req := new(dto.AlleleAddReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, dto.NewErrResult(cst.JSON_ERROR, err.Error()))
		return
	}
	result := a.alleleSvc.Add(*req, c.GetInt64("user"))
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
	result := a.alleleSvc.AlleleAllSearch(*req)
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
	result := a.alleleSvc.Update(*req, c.GetInt64("user"))
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
	result := a.alleleSvc.Delete(*req, c.GetInt64("user"))
	c.JSON(http.StatusOK, result)
	return
}
