package service

import (
	"github.com/linpanic/biology-server/caches"
	"github.com/linpanic/biology-server/dao"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func InitNumber() {
	number := dao.GetMaxStrainNumber()
	if number == "" {
		caches.InitNumber(0)
		return
	}
	number = strings.TrimLeft(number, "#")
	formatInt, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	caches.InitNumber(formatInt + 1)
}
