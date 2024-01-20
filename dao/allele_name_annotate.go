package dao

import (
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SelectAlleleNameAnnotate(alleleNameId int64, annotate string) []model.AlleleNameAnnotate {
	var result []model.AlleleNameAnnotate
	tx := db.DbLink.Model(&model.AlleleNameAnnotate{})
	if alleleNameId != 0 {
		tx = tx.Where(model.AlleleNameAnnotate{AlleleNameId: alleleNameId})
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

func CreateAlleleNameAnnotate(dbLink *gorm.DB, alleleNameId int64, annotate []string, creator, createTime int64) error {
	var data []model.AlleleNameAnnotate
	for _, v := range annotate {
		if v == "" {
			continue
		}
		data = append(data, model.AlleleNameAnnotate{
			AlleleNameId: alleleNameId,
			Annotate:     v,
			CreatorId:    creator,
			CreateTime:   createTime,
		})
	}
	err := dbLink.Create(&data).Error
	return err
}

func DeleteAlleleNameAnnotate(dbLink *gorm.DB, alleleNameId int64) error {
	err := dbLink.Delete(&model.AlleleNameAnnotate{}, "allele_name_id = ?", alleleNameId).Error
	return err
}

func DeleteAlleleNameAnnotateByIds(dbLink *gorm.DB, ids []int64) error {
	err := dbLink.Delete(&model.AlleleNameAnnotate{}, ids).Error
	return err
}

func UpdateAlleleNameAnnotate(dbLink *gorm.DB, id int64, annotate string, updateTime int64) error {
	err := dbLink.Model(&model.AlleleNameAnnotate{Id: id}).Updates(model.AlleleNameAnnotate{Annotate: annotate, UpdateTime: updateTime}).Error
	return err
}
