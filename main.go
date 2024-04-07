package main

import (
	"flag"
	"github.com/linpanic/biology-server/caches"
	"github.com/linpanic/biology-server/config"
	"github.com/linpanic/biology-server/cst"
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/logs"
)

func init() {
	logs.LogInit()
	db.DbInit()
}

var cfg = flag.String("f", "./config.json", "config file path")

func main() {
	flag.Parse()
	c := config.LoadConfig(*cfg)
	if c == nil {
		caches.JWTKey = []byte(cst.DEFAULT_JWT_KEY)
		caches.JWTTime = cst.DEFAULT_JWT_TIME
	}
	if c.JWTKey == "" {
		caches.JWTKey = []byte(cst.DEFAULT_JWT_KEY)
	}

	if c.JWTTime == 0 {
		caches.JWTTime = cst.DEFAULT_JWT_TIME
	}

}
