package api

import (
	"github.com/gin-gonic/gin"
	"github.com/linpanic/biology-server/dto"
	"github.com/linpanic/biology-server/server/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var (
	StrainApi StrainAPI
)

type StrainAPI struct {
	service.StrainService
	service.AlleleService
}

func (s *StrainAPI) GetNumber(c *gin.Context) {
	req := new(dto.StrainNumberReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, dto.NewErrResult(err.Error()))
		return
	}
	result := s.StrainService.GetNumber(*req)
	c.JSON(http.StatusOK, result)
	return
}

func (s *StrainAPI) Add(c *gin.Context) {
	req := new(dto.StrainAddReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, dto.NewErrResult(err.Error()))
		return
	}
	result := s.StrainService.Add(*req, c.GetInt64("user"))
	c.JSON(http.StatusOK, result)
	return
}

func (s *StrainAPI) List(c *gin.Context) {
	req := new(dto.StrainListReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, dto.NewErrResult(err.Error()))
		return
	}
	result := s.StrainService.List(*req)
	c.JSON(http.StatusOK, result)
	return
}

func (s *StrainAPI) StrainUpdate(c *gin.Context) {
	req := new(dto.StrainUpdateReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, dto.NewErrResult(err.Error()))
		return
	}
	result := s.StrainService.Update(*req, c.GetInt64("user"))
	c.JSON(http.StatusOK, result)
	return
}

func (s *StrainAPI) AlleleUpdate(c *gin.Context) {
	req := new(dto.AlleleUpdateReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, dto.NewErrResult(err.Error()))
		return
	}
	result := s.AlleleService.Update(*req, c.GetInt64("user"))
	c.JSON(http.StatusOK, result)
	return
}

func (s *StrainAPI) AlleleSearch(c *gin.Context) {
	req := new(dto.AlleleSearchReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, dto.NewErrResult(err.Error()))
		return
	}
	result := s.AlleleService.AlleleSearch(*req)
	c.JSON(http.StatusOK, result)
	return
}

func (s *StrainAPI) StrainDelete(c *gin.Context) {
	req := new(dto.StrainDelReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, dto.NewErrResult(err.Error()))
		return
	}
	result := s.StrainService.Delete(*req)
	c.JSON(http.StatusOK, result)
	return
}
