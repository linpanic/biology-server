package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/server/api"
	"github.com/linpanic/biology-server/server/middleware"
	"github.com/linpanic/biology-server/server/service"
	log "github.com/sirupsen/logrus"
)

func WebApiRun(port int64) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	dbLink := db.NewDB()
	alleleSvc := service.NewAlleleService(dbLink)
	strainSvc := service.NewStrainService(dbLink)

	//登录，注册不需要验证
	router.POST("/register", api.UserApi.Register)
	router.POST("/login", api.UserApi.Login)

	//鉴权
	router.POST("/valid", api.OauthApi.Valid)

	//需要验证

	strainApi := api.NewStrainApi(strainSvc, alleleSvc)

	biology := router.Group("/biology")
	biology.POST("/strain_list", strainApi.List)

	authGroup := biology.Group("/")
	authGroup.Use(middleware.JWTAndCasbinAuth())

	//品系
	authGroup.POST("/get_number", strainApi.GetNumber)
	authGroup.POST("/strain_add", strainApi.Add)
	authGroup.POST("/strain_update", strainApi.StrainUpdate)
	authGroup.POST("/strain_delete", strainApi.StrainDelete)

	//基因
	authGroup.POST("/allele_search", strainApi.AlleleSearch) //搜素列表信息

	alleleAPI := api.NewAlleleAPI(alleleSvc)
	{
		router.POST("/allele_list", alleleAPI.List)
		authGroup.POST("/allele_update", alleleAPI.Update)
		authGroup.POST("/allele_add", alleleAPI.Add)
		authGroup.POST("/allele_delete", alleleAPI.Delete)
	}

	//authGroup.POST("/delete_strain")
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}
