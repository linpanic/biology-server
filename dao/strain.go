package dao

import (
	"errors"
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// 精准搜索
func SelectOneStrain(dbLink *gorm.DB, id, number int64) *model.Strain {
	result := new(model.Strain)
	result.Id = id
	tx := dbLink.Model(&model.Strain{})
	if number != 0 {
		tx = tx.Where("number = ?", number)
	}
	err := tx.First(result).Error
	if err != nil {
		log.Error(err)
		return nil
	}
	return result
}

// 模糊搜索
func SelectStrain(dbLink *gorm.DB, name string, pageNo, pageSize int) []model.Strain {
	var result []model.Strain
	tx := dbLink.Model(&model.Strain{})

	if name != "" {
		tx = tx.Where("strain_name like ?", name)
	}
	//分页
	offset := (pageNo - 1) * pageSize
	err := tx.Find(&result).Offset(offset).Limit(pageSize).Error
	if err != nil {
		log.Error(err)
		return nil
	}
	return result
}

// 精准搜索
func SelectStrainByName(dbLink *gorm.DB, name string) *model.Strain {
	result := new(model.Strain)
	tx := dbLink.Model(&model.Strain{})

	if name != "" {
		tx = tx.Where("strain_name = ?", name)
	}

	err := tx.Take(result).Error
	if err != nil {
		log.Error(err)
		return nil
	}

	return result
}

func SelectStrainByNum(num string) *model.Strain {
	result := new(model.Strain)
	tx := db.DbLink.Model(&model.Strain{})

	if num != "" {
		tx = tx.Where("number = ?", num)
	}

	err := tx.Take(result).Error
	if err != nil {
		log.Error(err)
		return nil
	}

	return result
}

// 精准搜索
func SelectStrainByIds(ids []int64) []model.Strain {
	var result []model.Strain
	err := db.DbLink.Model(&model.Strain{}).Where("id in ?", ids).Find(&result).Error
	if err != nil {
		log.Error(err)
		return nil
	}
	return result
}

// 新增
func CreateStrain(dbLink *gorm.DB, num, name string, creator, createTime int64) (*model.Strain, error) {
	data := new(model.Strain)
	data.Number = num
	data.StrainName = name
	data.CreatorId = creator
	data.CreateTime = createTime
	err := dbLink.Create(data).Error
	return data, err
}

// 修改
func UpdateStrain(dbLink *gorm.DB, id int64, name string, updateTime int64) error {
	err := dbLink.Model(&model.Strain{Id: id}).Updates(model.Strain{StrainName: name, UpdateTime: updateTime}).Error
	return err
}

// 通过序列号修改
func UpdateStrainByNum(dbLink *gorm.DB, num string, name string, updateTime int64) error {
	err := dbLink.Model(&model.Strain{}).
		Where(&model.Strain{Number: num}).
		Updates(model.Strain{StrainName: name, UpdateTime: updateTime}).Error
	return err
}

// 删除一个
func DeleteOneStrain(dbLink *gorm.DB, id int64) error {
	err := dbLink.Delete(&model.Strain{Id: id}).Error
	return err
}

// 删除多个
func DeleteStrain(dbLink *gorm.DB, ids []int64) error {
	err := dbLink.Delete(&model.Strain{}, ids).Error
	return err
}

func GetMaxStrainNumber() string {
	result := new(model.Strain)
	err := db.DbLink.Model(model.Strain{}).Order("number desc").Take(result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ""
		}
		panic(err)
	}
	return result.Number
}
