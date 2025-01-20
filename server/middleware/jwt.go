package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/linpanic/biology-server/cst"
	"github.com/linpanic/biology-server/dao"
	"github.com/linpanic/biology-server/dto"
	"github.com/linpanic/biology-server/permission"
	"github.com/linpanic/biology-server/utils"
	"net/http"
)

func JWTAndCasbinAuth() func(*gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			c.JSON(http.StatusOK, dto.NewErrResult(cst.NO_LOGIN, "未登录或非法访问"))
			c.Abort()
			return
		}

		userId, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, dto.LoginErrorResult)
			c.Abort()
			return
		}

		user := dao.SelectUserById(userId)
		if user == nil {
			c.JSON(http.StatusOK, dto.NewErrResult(cst.USER_NOT_EXIST, "找不到该用户"))
			c.Abort()
			return
		}
		b, _ := permission.Ef.Enforce(user.UserName, c.Request.URL.Path, "read")
		if b {
			c.JSON(http.StatusOK, dto.NewErrResult(cst.NO_ACCESS, "未授权用户"))
			c.Abort()
			return
		}

		c.Set("user", userId)
		c.Next()
	}
}
