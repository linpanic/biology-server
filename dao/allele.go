package dao

import (
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/dto"
	"github.com/linpanic/biology-server/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SelectAllele(strainId int64) []model.Allele {
	var result []model.Allele
	err := db.DbLink.Model(&model.Allele{}).Where(model.Allele{StrainId: strainId}).Find(&result).Error
	if err != nil {
		log.Error(err)
		return nil
	}
	return result
}

func SelectAlleleById(id int64) *model.Allele {
	result := new(model.Allele)
	err := db.DbLink.Model(&model.Allele{Id: id}).First(result).Error
	if err != nil {
		return nil
	}
	return result
}

func SelectAlleleByName(name string) []model.Allele {
	var result []model.Allele
	err := db.DbLink.Model(&model.Allele{}).Select("id,name,genome,serial").Where("name like ?", "%"+name+"%").Find(&result).Error
	if err != nil {
		log.Error(err)
		return nil
	}
	return result
}

func CreateAllele(dbLink *gorm.DB, strainId int64, req dto.Allele, creator, createTime int64) (*model.Allele, error) {
	data := new(model.Allele)
	data.StrainId = strainId
	data.Name = req.Name
	data.Genome = req.Genome
	data.Serial = req.Serial
	data.CreatorId = creator
	data.CreateTime = createTime
	err := dbLink.Create(&data).Error
	return data, err
}

func CreateAlleles(dbLink *gorm.DB, strainId int64, datas []dto.Allele, creator, createTime int64) ([]model.Allele, error) {
	var adds []model.Allele
	for _, v := range datas {
		var data model.Allele
		data.StrainId = strainId
		data.Name = v.Name
		data.Genome = v.Genome
		data.Serial = v.Serial
		data.CreatorId = creator
		data.CreateTime = createTime
		adds = append(adds, data)
	}
	if len(adds) == 0 {
		return nil, nil
	}
	err := dbLink.Create(&adds).Error
	return adds, err
}

func UpdateAllele(dbLink *gorm.DB, id int64, name, genome, serial string, updateTime int64) error {
	err := dbLink.Model(&model.Allele{Id: id}).Select("name", "genome", "serial").
		Updates(model.Allele{
			Name:       name,
			Genome:     genome,
			Serial:     serial,
			UpdateTime: updateTime,
		}).Error
	return err
}

func DeleteAllele(dbLink *gorm.DB, strainId int64) error {
	err := dbLink.Debug().Delete(&model.Allele{}, "strain_id = ?", strainId).Error
	return err
}

func DeleteAlleleByIds(dbLink *gorm.DB, ids []int64) error {
	err := dbLink.Delete(&model.Allele{}, ids).Error
	return err
}
