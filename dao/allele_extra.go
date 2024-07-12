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

func CreateAlleleExtra(dbLink *gorm.DB, alleleId int64, extra []dto.ExtraInfo, creator, createTime int64) error {
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
			AlleleId:   alleleId,
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

func CreateAllelesExtra(dbLink *gorm.DB, req []dto.Allele, creator, createTime int64) error {
	var data []model.AlleleExtra
	for _, v := range req {
		if v.Id == 0 {
			continue
		}

		for _, v2 := range v.Extra {
			if v2.ExtraKey == "" && v2.ExtraVal == "" {
				continue
			}
			if v2.ExtraKey == "" {
				v2.ExtraKey = " "
			}
			if v2.ExtraVal == "" {
				v2.ExtraVal = " "
			}
			data = append(data, model.AlleleExtra{
				AlleleId:   v2.Id,
				ExtraKey:   v2.ExtraKey,
				ExtraValue: v2.ExtraVal,
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

func DeleteAlleleExtra(dbLink *gorm.DB, alleleId int64) error {
	err := dbLink.Delete(&model.AlleleExtra{}, "allele_id = ?", alleleId).Error
	return err
}

func DeleteAlleleExtras(dbLink *gorm.DB, alleleIds []int64) error {
	err := dbLink.Delete(&model.AlleleExtra{}, "allele_id in ?", alleleIds).Error
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
