package api

import (
	"github.com/gin-gonic/gin"
	"github.com/linpanic/biology-server/cst"
	"github.com/linpanic/biology-server/dao"
	"github.com/linpanic/biology-server/dto"
	"github.com/linpanic/biology-server/utils"
	"net/http"
)

var (
	OauthApi OauthAPI
)

type OauthAPI struct{}

func (o *OauthAPI) Valid(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	if token == "" {
		c.JSON(http.StatusOK, dto.NoTokenResult())
		c.Abort()
		return
	}

	userId, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, dto.TokenErrorResult())
		c.Abort()
		return
	}

	user := dao.SelectUserById(userId)
	if user == nil {
		c.JSON(http.StatusOK, dto.NewErrResult(cst.USER_NOT_EXIST, "找不到该用户"))
		c.Abort()
		return
	}
	//c.Set("user", userId)
	//c.Next()
	c.JSON(http.StatusOK, dto.NewOKResult(nil))
}
