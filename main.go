package main

import (
	"flag"
	"github.com/linpanic/biology-server/caches"
	"github.com/linpanic/biology-server/config"
	"github.com/linpanic/biology-server/cst"
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/logs"
	"github.com/linpanic/biology-server/permission"
	"github.com/linpanic/biology-server/server/router"
	"github.com/linpanic/biology-server/server/service"
	log "github.com/sirupsen/logrus"
)

func init() {
	//初始化日志
	logs.LogInit()

	//初始化数据库链接
	db.DbInit()

	//初始化品系序列号
	service.InitNumber()

	caches.InitStrainAlleleField()

	permission.InitCasbin()

}

var cfg = flag.String("f", "./config.json", "config file path")

func main() {
	flag.Parse()
	c := config.LoadConfig(*cfg)

	if c.JWTKey == "" {
		caches.JWTKey = []byte(cst.DEFAULT_JWT_KEY)
	}

	if c.JWTTime == 0 {
		caches.JWTTime = cst.DEFAULT_JWT_TIME
	}
	if c.Port == 0 {
		c.Port = 10080
	}

	router.WebApiRun(c.Port)
	log.Info("project from https://github.com/linpanic/biology-server")
	select {}
}
