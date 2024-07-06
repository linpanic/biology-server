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
