package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReturnOK(c *gin.Context, result any) {
	c.JSON(http.StatusOK, result)
}
