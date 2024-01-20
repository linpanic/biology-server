package dao

import (
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// 精准搜索
func SelectOneAlleleNameById(id int64) *model.AlleleName {
	result := new(model.AlleleName)
	result.Id = id

	err := db.DbLink.Model(&model.AlleleName{}).First(result).Error

	if err != nil {
		log.Error(err)
		return nil
	}
	return result
}

// 模糊搜索
func SelectAlleleName(strainId int64, name string) []model.AlleleName {
	var result []model.AlleleName
	tx := db.DbLink.Model(&model.AlleleName{})
	if strainId != 0 {
		tx = tx.Where("strain_id = ?", strainId)
	}
	if name != "" {
		tx = tx.Where("name like ?", name)
	}
	err := tx.Find(&result).Error
	if err != nil {
		log.Error(err)
		return nil
	}
	return result
}

// 精准搜索
func SelectAlleleNameByIds(ids []int64) []model.AlleleName {
	var result []model.AlleleName
	err := db.DbLink.Model(&model.AlleleName{}).Where("id in ?", ids).Find(&result).Error
	if err != nil {
		log.Error(err)
		return nil
	}
	return result
}

// 新增
func CreateAlleleName(dbLink *gorm.DB, strainId int64, name string, creator, createTime int64) error {
	data := new(model.AlleleName)
	data.StrainId = strainId
	data.AlleleName = name
	data.CreatorId = creator
	data.CreateTime = createTime
	err := dbLink.Create(data).Error
	return err
}

// 修改
func UpdateAlleleName(dbLink *gorm.DB, id int64, name string, updateTime int64) error {
	err := dbLink.Model(&model.AlleleName{Id: id}).Updates(model.AlleleName{AlleleName: name, UpdateTime: updateTime}).Error
	return err
}

// 删除一个
func DeleteOneAlleleName(dbLink *gorm.DB, id int64) error {
	err := dbLink.Delete(&model.AlleleName{Id: id}).Error
	return err
}

// 删除多个
func DeleteAlleleName(dbLink *gorm.DB, ids []int64) error {
	err := dbLink.Delete(&model.AlleleName{}, ids).Error
	return err
}

// 通过品系ID删除多个
func DeleteAlleleNameByStrainId(dbLink *gorm.DB, strainId int64) error {
	err := dbLink.Delete(&model.AlleleName{}, "strain_id = ?", strainId).Error
	return err
}
