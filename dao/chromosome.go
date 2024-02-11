package dao

import (
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SelectChromsome(genomeId int64) []model.Chromosome {
	var result []model.Chromosome
	err := db.DbLink.Model(&model.Chromosome{}).Where(model.Chromosome{GenomeId: genomeId}).Find(&result).Error
	if err != nil {
		log.Error(err)
		return nil
	}
	return result
}

func CreateChromsome(dbLink *gorm.DB, genomeId int64, serials []string, createTime int64) error {
	var data []model.Chromosome
	for _, v := range serials {
		if v == "" {
			continue
		}
		data = append(data, model.Chromosome{
			GenomeId:   genomeId,
			Serial:     v,
			CreateTime: createTime,
		})
	}
	err := dbLink.Create(&data).Error
	return err
}

func UpdateChromsome(dbLink *gorm.DB, id int64, serial string, updateTime int64) error {
	err := dbLink.Model(&model.Chromosome{Id: id}).Updates(model.Chromosome{Serial: serial, UpdateTime: updateTime}).Error
	return err
}

func DeleteChromsome(dbLink *gorm.DB, genomeId int64) error {
	err := dbLink.Delete(&model.Chromosome{}, "genome_id = ?", genomeId).Error
	return err
}

func DeleteChromsomeByIds(dbLink *gorm.DB, ids []int64) error {
	err := dbLink.Delete(&model.Chromosome{}, ids).Error
	return err
}
