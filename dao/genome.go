package dao

import (
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SelectOneGenome(alleleId int64, name string) *model.Genome {
	result := new(model.Genome)
	result.AllelNameId = alleleId

	tx := db.DbLink.Model(&model.Genome{})
	if name != "" {
		tx = tx.Where("genome_name like ?", name)
	}

	err := tx.First(result).Error
	if err != nil {
		log.Error(err)
		return nil
	}
	return result
}

func SelectGenome(pageNo, pageSize int) []model.Genome {
	if pageNo == 0 {
		pageNo = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}

	var result []model.Genome

	//分页
	offset := (pageNo - 1) * pageSize
	err := db.DbLink.Model(&model.Genome{}).Offset(offset).Limit(pageSize).Find(&result).Error
	if err != nil {
		log.Error(err)
		return nil
	}
	return result
}

func CreateGenome(dbLink *gorm.DB, alleleId int64, name string, creator, createTime int64) (*model.Genome, error) {
	data := new(model.Genome)
	data.AllelNameId = alleleId
	data.CreateTime = createTime
	data.GenomeName = name
	data.CreatorId = creator
	err := dbLink.Create(data).Error
	return data, err
}

func UpdateGenome(dbLink *gorm.DB, id int64, name string, ut int64) error {
	err := dbLink.Model(&model.Genome{Id: id}).Updates(model.Genome{GenomeName: name, UpdateTime: ut}).Error
	return err
}
