package main

import (
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/logs"
)

func init() {
	logs.LogInit()
	db.DbInit()
}

func main() {

}
