package api

import (
	"github.com/gin-gonic/gin"
	"github.com/linpanic/biology-server/dto"
	"github.com/linpanic/biology-server/server/service"
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
		ReturnJSONError(c, err)
		return
	}
	result := a.alleleSvc.Add(*req, c.GetInt64("user"))
	ReturnOK(c, result)
	return
}

func (a *AlleleAPI) List(c *gin.Context) {
	req := new(dto.AlleleListReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		ReturnJSONError(c, err)
		return
	}
	result := a.alleleSvc.AlleleAllSearch(*req)
	ReturnOK(c, result)
	return
}

func (a *AlleleAPI) Update(c *gin.Context) {
	req := new(dto.AlleleUpdateReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		ReturnJSONError(c, err)
		return
	}
	result := a.alleleSvc.Update(*req, c.GetInt64("user"))
	ReturnOK(c, result)
	return
}

func (a *AlleleAPI) Delete(c *gin.Context) {
	req := new(dto.AlleleDelReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		ReturnJSONError(c, err)
		return
	}
	result := a.alleleSvc.Delete(*req, c.GetInt64("user"))
	ReturnOK(c, result)
	return
}
