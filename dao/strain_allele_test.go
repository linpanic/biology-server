package dao

import (
	"encoding/json"
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/logs"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestSelectStrainAndAllele(t *testing.T) {
	logs.LogInit()
	db.DbInit()

	allele, i := SelectStrainAndAllele("华为", "", "", 1, 10)
	marshal, _ := json.Marshal(allele)
	log.Info(string(marshal))
	log.Info(i)
}

type Student struct {
	Name string
}
