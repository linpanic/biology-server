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

func CreateAllele(dbLink *gorm.DB, strainId int64, req dto.Allele, creator, createTime int64) (*model.Allele, error) {
	data := new(model.Allele)
	data.StrainId = strainId
	data.Name = req.AlleleName
	data.Genome = req.GenomeName
	data.Serial = req.Serial
	data.CreatorId = creator
	data.CreateTime = createTime
	err := dbLink.Create(&data).Error
	return data, err
}

func CreateAlleles(dbLink *gorm.DB, strainId int64, datas []dto.Allele, creator, createTime int64) error {
	var adds []model.Allele
	for _, v := range datas {
		var data model.Allele
		data.StrainId = strainId
		data.Name = v.AlleleName
		data.Genome = v.GenomeName
		data.Serial = v.Serial
		data.CreatorId = creator
		data.CreateTime = createTime
		adds = append(adds, data)
	}
	err := dbLink.Create(&adds).Error
	return err
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
	err := dbLink.Delete(&model.Allele{}, "strain_id = ?", strainId).Error
	return err
}

func DeleteAlleleByIds(dbLink *gorm.DB, ids []int64) error {
	err := dbLink.Delete(&model.Allele{}, ids).Error
	return err
}
