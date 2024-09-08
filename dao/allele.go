package dao

import (
	"fmt"
	"github.com/linpanic/biology-server/cst"
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/dto"
	"github.com/linpanic/biology-server/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SelectAlleleByAll(kw, field, order string, pageNo, pageSize int) ([]model.AlleleAll, int64) {
	sql := cst.ALLELE_LIST_SQL
	if kw != "" {
		sql += cst.ALLELE_WHERE_SQL
	}

	countSql := fmt.Sprintf(cst.STRAIN_COUNT_SQL, sql)

	if field != "" {
		sql += " order by " + field + " "
		if order != "" {
			sql += order
		} else {
			sql += " desc"
		}
	}

	var err error
	var count int64
	offset := (pageNo - 1) * pageSize

	var result []model.AlleleAll
	sql += " LIMIT ? OFFSET ?"
	if kw != "" {
		kw = "%" + kw + "%"
		err = db.DbLink.Debug().Raw(sql, kw, kw, kw, kw, kw, kw, kw, kw, kw, kw, kw, kw, pageSize, offset).Scan(&result).Error
		if err != nil {
			log.Error(err)
			return nil, 0
		}
		err = db.DbLink.Raw(countSql, kw, kw, kw, kw, kw, kw, kw, kw, kw, kw, kw, kw).Count(&count).Error
	} else {
		err = db.DbLink.Debug().Raw(sql, pageSize, offset).Scan(&result).Error
		if err != nil {
			log.Error(err)
			return nil, 0
		}
		err = db.DbLink.Debug().Raw(countSql).Count(&count).Error
	}
	if err != nil {
		log.Error(err)
		return nil, 0
	}
	return result, count
}

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
