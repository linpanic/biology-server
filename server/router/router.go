package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/linpanic/biology-server/server/api"
	"github.com/linpanic/biology-server/server/middleware"
	log "github.com/sirupsen/logrus"
)

func WebApiRun(port int64) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//登录，注册不需要验证
	router.POST("/register", api.UserApi.Register)
	router.POST("/login", api.UserApi.Login)

	//鉴权
	router.POST("/valid", api.OauthApi.Valid)

	//需要验证
	biology := router.Group("/biology")
	biology.POST("/strain_list", api.StrainApi.List)

	authGroup := biology.Group("/")
	authGroup.Use(middleware.JWTAuth())

	//品系
	authGroup.POST("/get_number", api.StrainApi.GetNumber)
	authGroup.POST("/strain_add", api.StrainApi.Add)
	authGroup.POST("/strain_update", api.StrainApi.StrainUpdate)
	authGroup.POST("/allele_update", api.StrainApi.AlleleUpdate)
	authGroup.POST("/strain_delete", api.StrainApi.StrainDelete)

	//基因
	authGroup.POST("/allele_search", api.StrainApi.AlleleSearch)

	//authGroup.POST("/delete_strain")
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}
