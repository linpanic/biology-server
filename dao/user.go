package dao

import (
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func CreateUser(dbLink *gorm.DB, name, pw string, createTime int64) error {
	err := dbLink.Create(&model.User{
		UserName:   name,
		Password:   pw,
		CreateTime: createTime,
	}).Error
	return err
}

func SelectOneUser(name, pw string) *model.User {
	result := new(model.User)
	err := db.DbLink.Model(&model.User{}).Where(model.User{UserName: name, Password: pw}).First(result).Error
	if err != nil {
		log.Error(err)
		return nil
	}
	return result
}

func SelectUserCount(name string) int64 {
	result := new(int64)
	err := db.DbLink.Model(&model.User{}).Where(model.User{UserName: name}).Count(result).Error
	if err != nil {
		return 0
	}
	return *result
}

func SelectUserById(id int64) *model.User {
	result := new(model.User)
	result.Id = id
	err := db.DbLink.Model(&model.User{}).First(result).Error
	if err != nil {
		log.Error(err)
		return nil
	}
	return result
}

func UpdateUser(dbLink *gorm.DB, id int64, pw string, upt int64) error {
	err := dbLink.Model(&model.User{Id: id}).
		Updates(model.User{Password: pw, UpdateTime: upt}).
		Error
	return err
}

func DeleteUser(dbLink *gorm.DB, id int64) error {
	err := dbLink.Delete(&model.User{Id: id}).Error
	return err
}
