package router

import (
	"github.com/gin-gonic/gin"
	"github.com/linpanic/biology-server/server/api"
	"github.com/linpanic/biology-server/server/middleware"
)

func WebApiRun() {
	router := gin.Default()

	//登录，注册不需要验证
	router.POST("/register", api.UserApi.Register)
	router.POST("/login", api.UserApi.Login)

	//需要验证
	authGroup := router.Group("/biology")
	authGroup.Use(middleware.JWTAuth())

	//品系
	authGroup.POST("/strain_list,")
	authGroup.POST("/add_strain,")
	authGroup.POST("/update_strain,")
	authGroup.POST("/delete_strain,")

}
