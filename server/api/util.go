package api

import (
	"github.com/gin-gonic/gin"
	"github.com/linpanic/biology-server/cst"
	"github.com/linpanic/biology-server/dto"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func ReturnOK(c *gin.Context, result any) {
	c.JSON(http.StatusOK, result)
}

func ReturnJSONError(c *gin.Context, err error) {
	log.Error(err)
	c.JSON(http.StatusOK, dto.NewErrResult(cst.JSON_ERROR, err.Error()))
}
