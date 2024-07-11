package db

import (
	"github.com/linpanic/biology-server/model"
	"github.com/linpanic/biology-server/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DbLink *gorm.DB
	DBPath = "C:\\db\\biology.db"
	DBDir  = "C:\\db"
)

func DbInit() {
	exist := utils.CheckFileExist(DBDir)
	if !exist {
		utils.CreateDir(DBDir)
	}
	exist = utils.CheckFileExist(DBPath)
	var err error
	DbLink, err = gorm.Open(sqlite.Open(DBPath), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	err = DbLink.AutoMigrate(&model.Allele{}, &model.AlleleAnnotate{},
		&model.AlleleExtra{}, &model.ShortName{}, &model.Strain{},
		&model.StrainAnnotate{}, &model.StrainExtra{}, &model.User{})
	if err != nil {
		log.Fatal(err)
	}
}
