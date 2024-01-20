package dao

import (
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/dto"
	"github.com/linpanic/biology-server/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SelectAlleleNameExtra(alleleNameId int64, extra string) []model.AlleleNameExtra {
	var result []model.AlleleNameExtra
	tx := db.DbLink.Model(&model.AlleleNameExtra{})
	if alleleNameId != 0 {
		tx = tx.Where(model.AlleleNameExtra{AlleleNameId: alleleNameId})
	}
	if extra != "" {
		tx = tx.Where("extra like ?", extra)
	}
	err := tx.Find(&result).Error
	if err != nil {
		log.Error(err)
		return nil
	}
	return result
}

func CreateAlleleNameExtra(dbLink *gorm.DB, alleleNameId int64, extra []dto.ExtraInfo, creator, createTime int64) error {
	var data []model.AlleleNameExtra
	for _, v := range extra {
		if v.ExtraKey == "" && v.ExtraVal == "" {
			continue
		}
		data = append(data, model.AlleleNameExtra{
			AlleleNameId: alleleNameId,
			ExtraKey:     v.ExtraKey,
			ExtraValue:   v.ExtraVal,
			CreatorId:    creator,
			CreateTime:   createTime,
		})
	}
	if len(data) == 0 {
		return nil
	}
	err := dbLink.Create(&data).Error
	return err
}

func DeleteAlleleNameExtra(dbLink *gorm.DB, alleleNameId int64) error {
	err := dbLink.Delete(&model.AlleleNameExtra{}, "allele_name_id = ?", alleleNameId).Error
	return err
}

func DeleteAlleleNameExtraByIds(dbLink *gorm.DB, ids []int64) error {
	err := dbLink.Delete(&model.AlleleNameExtra{}, ids).Error
	return err
}

func UpdateAlleleNameExtra(dbLink *gorm.DB, id int64, extra dto.ExtraInfo, updateTime int64) error {
	err := dbLink.Model(&model.AlleleNameExtra{Id: id}).
		Updates(model.AlleleNameExtra{ExtraKey: extra.ExtraKey, ExtraValue: extra.ExtraVal, UpdateTime: updateTime}).
		Error
	return err
}
