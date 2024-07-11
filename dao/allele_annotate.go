package dao

import (
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SelectAlleleAnnotate(alleleNameId int64, annotate string) []model.AlleleAnnotate {
	var result []model.AlleleAnnotate
	tx := db.DbLink.Model(&model.AlleleAnnotate{})
	if alleleNameId != 0 {
		tx = tx.Where(model.AlleleAnnotate{AlleleId: alleleNameId})
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

func CreateAlleleAnnotate(dbLink *gorm.DB, alleleNameId int64, annotate []string, creator, createTime int64) error {
	var data []model.AlleleAnnotate
	for _, v := range annotate {
		if v == "" {
			continue
		}
		data = append(data, model.AlleleAnnotate{
			AlleleId:   alleleNameId,
			Annotate:   v,
			CreatorId:  creator,
			CreateTime: createTime,
		})
	}
	err := dbLink.Create(&data).Error
	return err
}

func DeleteAlleleAnnotate(dbLink *gorm.DB, alleleNameId int64) error {
	err := dbLink.Delete(&model.AlleleAnnotate{}, "allele_id = ?", alleleNameId).Error
	return err
}

func DeleteAlleleAnnotateByIds(dbLink *gorm.DB, ids []int64) error {
	err := dbLink.Delete(&model.AlleleAnnotate{}, ids).Error
	return err
}

func UpdateAlleleAnnotate(dbLink *gorm.DB, id int64, annotate string, updateTime int64) error {
	err := dbLink.Model(&model.AlleleAnnotate{Id: id}).Updates(model.AlleleAnnotate{Annotate: annotate, UpdateTime: updateTime}).Error
	return err
}
