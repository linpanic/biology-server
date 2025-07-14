package api

import (
	"github.com/gin-gonic/gin"
	"github.com/linpanic/biology-server/dto"
	"github.com/linpanic/biology-server/server/service"
)

//var (
//	StrainApi StrainAPI
//)

func NewStrainApi(strainSvc *service.StrainService, alleleSvc *service.AlleleService) *StrainAPI {
	return &StrainAPI{
		strainSvc:  strainSvc,
		allleleSvc: alleleSvc,
	}
}

type StrainAPI struct {
	strainSvc  *service.StrainService
	allleleSvc *service.AlleleService
}

func (s *StrainAPI) GetNumber(c *gin.Context) {
	req := new(dto.StrainNumberReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		ReturnJSONError(c, err)
		return
	}
	result := s.strainSvc.GetNumber(*req)
	ReturnOK(c, result)
	return
}

func (s *StrainAPI) Add(c *gin.Context) {
	req := new(dto.StrainAddReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		ReturnJSONError(c, err)
		return
	}
	result := s.strainSvc.Add(*req, c.GetInt64("user"))
	ReturnOK(c, result)
	return
}

func (s *StrainAPI) List(c *gin.Context) {
	req := new(dto.StrainListReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		ReturnJSONError(c, err)
		return
	}
	result := s.strainSvc.List(*req)
	ReturnOK(c, result)
	return
}

func (s *StrainAPI) StrainUpdate(c *gin.Context) {
	req := new(dto.StrainUpdateReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		ReturnJSONError(c, err)
		return
	}
	result := s.strainSvc.Update(*req, c.GetInt64("user"))
	ReturnOK(c, result)
	return
}

func (s *StrainAPI) AlleleUpdate(c *gin.Context) {
	req := new(dto.AlleleUpdateReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		ReturnJSONError(c, err)
		return
	}
	result := s.allleleSvc.Update(*req, c.GetInt64("user"))
	ReturnOK(c, result)
	return
}

func (s *StrainAPI) AlleleSearch(c *gin.Context) {
	req := new(dto.AlleleSearchReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		ReturnJSONError(c, err)
		return
	}
	result := s.allleleSvc.AlleleSearch(*req)
	ReturnOK(c, result)
	return
}

func (s *StrainAPI) StrainDelete(c *gin.Context) {
	req := new(dto.StrainDelReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		ReturnJSONError(c, err)
		return
	}
	result := s.strainSvc.Delete(*req)
	ReturnOK(c, result)
	return
}
