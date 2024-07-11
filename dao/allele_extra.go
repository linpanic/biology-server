package dao

import (
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/dto"
	"github.com/linpanic/biology-server/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SelectAlleleExtra(alleleNameId int64, extra string) []model.AlleleExtra {
	var result []model.AlleleExtra
	tx := db.DbLink.Model(&model.AlleleExtra{})
	if alleleNameId != 0 {
		tx = tx.Where(model.AlleleExtra{AlleleId: alleleNameId})
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

func CreateAlleleExtra(dbLink *gorm.DB, alleleNameId int64, extra []dto.ExtraInfo, creator, createTime int64) error {
	var data []model.AlleleExtra
	for _, v := range extra {
		if v.ExtraKey == "" && v.ExtraVal == "" {
			continue
		}
		if v.ExtraKey == "" {
			v.ExtraKey = " "
		}
		if v.ExtraVal == "" {
			v.ExtraVal = " "
		}
		data = append(data, model.AlleleExtra{
			AlleleId:   alleleNameId,
			ExtraKey:   v.ExtraKey,
			ExtraValue: v.ExtraVal,
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

func DeleteAlleleExtra(dbLink *gorm.DB, alleleNameId int64) error {
	err := dbLink.Delete(&model.AlleleExtra{}, "allele_id = ?", alleleNameId).Error
	return err
}

func DeleteAlleleExtraByIds(dbLink *gorm.DB, ids []int64) error {
	err := dbLink.Delete(&model.AlleleExtra{}, ids).Error
	return err
}

func UpdateAlleleExtra(dbLink *gorm.DB, id int64, extra dto.ExtraInfo, updateTime int64) error {
	err := dbLink.Model(&model.AlleleExtra{Id: id}).
		Updates(model.AlleleExtra{ExtraKey: extra.ExtraKey, ExtraValue: extra.ExtraVal, UpdateTime: updateTime}).
		Error
	return err
}
