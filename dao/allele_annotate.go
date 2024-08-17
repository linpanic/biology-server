package dao

import (
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/dto"
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

func CreateAlleleAnnotate(dbLink *gorm.DB, alleleId int64, annotate []string, creator, createTime int64) error {
	var data []model.AlleleAnnotate
	for _, v := range annotate {
		if v == "" {
			continue
		}
		data = append(data, model.AlleleAnnotate{
			AlleleId:   alleleId,
			Annotate:   v,
			CreatorId:  creator,
			CreateTime: createTime,
		})
	}
	if len(data) == 0 {
		return nil
	}
	err := dbLink.Create(&data).Error
	return err
}

func CreateAllelesAnnotate(dbLink *gorm.DB, req []dto.Allele, creator, createTime int64) error {
	var data []model.AlleleAnnotate
	for _, v := range req {
		if v.Id == 0 {
			continue
		}
		for _, v2 := range v.Annotate {
			if v2 == "" {
				continue
			}
			data = append(data, model.AlleleAnnotate{
				AlleleId:   v.Id,
				Annotate:   v2,
				CreatorId:  creator,
				CreateTime: createTime,
			})
		}
	}
	if len(data) == 0 {
		return nil
	}
	err := dbLink.Create(&data).Error
	return err
}

func DeleteAlleleAnnotate(dbLink *gorm.DB, alleleId int64) error {
	err := dbLink.Delete(&model.AlleleAnnotate{}, "allele_id = ?", alleleId).Error
	return err
}

func DeleteAlleleAnnotates(dbLink *gorm.DB, alleleId []int64) error {
	err := dbLink.Delete(&model.AlleleAnnotate{}, "allele_id in ?", alleleId).Error
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
