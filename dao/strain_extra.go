package dao

import (
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/dto"
	"github.com/linpanic/biology-server/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SelectStrainExtra(strainId int64, extra string) []model.StrainExtra {
	var result []model.StrainExtra
	tx := db.DbLink.Model(&model.StrainExtra{})
	if strainId != 0 {
		tx = tx.Where(model.StrainExtra{StrainId: strainId})
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

func CreateStrainExtra(dbLink *gorm.DB, strainId int64, extra []dto.ExtraInfo, creator, createTime int64) error {
	var data []model.StrainExtra
	for _, v := range extra {
		if v.ExtraKey == "" && v.ExtraVal == "" {
			continue
		}
		data = append(data, model.StrainExtra{
			StrainId:   strainId,
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

func DeleteStrainExtra(dbLink *gorm.DB, strainId int64) error {
	err := dbLink.Delete(&model.StrainExtra{}, "strain_id = ?", strainId).Error
	return err
}

func DeleteStrainExtraByIds(dbLink *gorm.DB, ids []int64) error {
	err := dbLink.Delete(&model.StrainExtra{}, ids).Error
	return err
}

func UpdateStrainExtra(dbLink *gorm.DB, id int64, extra dto.ExtraInfo, updateTime int64) error {
	err := dbLink.Model(&model.StrainExtra{Id: id}).
		Updates(model.StrainExtra{ExtraKey: extra.ExtraKey, ExtraValue: extra.ExtraVal, UpdateTime: updateTime}).
		Error
	return err
}
