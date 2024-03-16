package dao

import (
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SelectStrainAnnotate(strainId int64, annotate string) []model.StrainAnnotate {
	var result []model.StrainAnnotate
	tx := db.DbLink.Model(&model.StrainAnnotate{})
	if strainId != 0 {
		tx = tx.Where(model.StrainAnnotate{StrainId: strainId})
	}
	if annotate != "" {
		tx = tx.Where("annotate like ?", annotate)
	}
	err := tx.Find(&result).Error
	if err != nil {
		log.Error(err)
		return nil
	}
	return result
}

func CreateStrainAnnotate(dbLink *gorm.DB, strainId int64, annotate []string, creator, createTime int64) error {
	var data []model.StrainAnnotate
	for _, v := range annotate {
		if v == "" {
			continue
		}
		data = append(data, model.StrainAnnotate{
			StrainId:   strainId,
			Annotate:   v,
			CreatorId:  creator,
			CreateTime: createTime,
		})
	}
	err := dbLink.Create(&data).Error
	return err
}

func DeleteStrainAnnotate(dbLink *gorm.DB, strainId int64) error {
	err := dbLink.Delete(&model.StrainAnnotate{}, "strain_id = ?", strainId).Error
	return err
}

func DeleteStrainAnnotateByIds(dbLink *gorm.DB, ids []int64) error {
	err := dbLink.Delete(&model.StrainAnnotate{}, ids).Error
	return err
}

func UpdateStrainAnnotate(dbLink *gorm.DB, id int64, annotate string, updateTime int64) error {
	err := dbLink.Model(&model.StrainAnnotate{Id: id}).Updates(model.StrainAnnotate{Annotate: annotate, UpdateTime: updateTime}).Error
	return err
}
