package dao

import (
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SelectShortName(strainId int64) []model.ShortName {
	var result []model.ShortName
	err := db.DbLink.Model(&model.ShortName{}).Where(model.ShortName{StrainId: strainId}).Find(&result).Error
	if err != nil {
		log.Error(err)
		return nil
	}
	return result
}

func CreateShortName(dbLink *gorm.DB, strainId int64, shortNames []string, creator, createTime int64) error {
	var data []model.ShortName
	for _, v := range shortNames {
		if v == "" {
			continue
		}
		data = append(data, model.ShortName{
			StrainId:   strainId,
			ShortName:  v,
			CreatorId:  creator,
			CreateTime: createTime,
		})
	}
	err := dbLink.Create(&data).Error
	return err
}

func UpdateShortName(dbLink *gorm.DB, id int64, name string, updateTime int64) error {
	err := dbLink.Model(&model.ShortName{Id: id}).Updates(model.ShortName{ShortName: name, UpdateTime: updateTime}).Error
	return err
}

func DeleteShortName(dbLink *gorm.DB, strainId int64) error {
	err := dbLink.Delete(&model.ShortName{}, "strain_id = ?", strainId).Error
	return err
}

func DeleteShortNameByIds(dbLink *gorm.DB, ids []int64) error {
	err := dbLink.Delete(&model.ShortName{}, ids).Error
	return err
}
