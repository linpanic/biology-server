package dao

import (
	"fmt"
	"github.com/linpanic/biology-server/cst"
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/model"
	log "github.com/sirupsen/logrus"
)

func SelectStrainAndAllele(kw, field, order string, pageNo, pageSize int) ([]model.StrainAllele, int64) {
	sql := cst.STRAIN_SQL
	if kw != "" {
		sql += cst.STRAIN_LIKE_SQL
	}
	sql += cst.STRAIN_END_SQL

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

	var result []model.StrainAllele
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

//
//func SelectStrainAndAllele(kw,field,order string,pageNo,pageSize int) ([]int64,int64) {
//	sql := cst.STRAIN_ALLELE_SQL
//
//	if kw != "" {
//		sql += cst.STRAIN_ALLELE_LIKE_SQL
//	}
//	sql += cst.STRAIN_ALLELE_END_SQL
//
//	if field != "" {
//		sql += "order by " + field
//		if order != "" {
//			sql += order
//		}else {
//			sql += "desc"
//		}
//	}
//
//	countSql := fmt.Sprintf(cst.STRAIN_ALLELE_COUNT_SQL,sql)
//
//	sql += " LIMIT ? OFFSET ?"
//
//	var strainIds []int64
//	var err error
//	var count int64
//	offset := (pageNo - 1) * pageSize
//
//	if kw != "" {
//		kw = "%"+kw+"%"
//		err = db.DbLink.Debug().Raw(sql, kw,kw, kw, kw, kw, kw, kw, kw, kw, kw, kw, kw,pageSize,offset).Scan(&strainIds).Error
//		if err != nil {
//			log.Error(err)
//			return nil,0
//		}
//		err = db.DbLink.Raw(countSql, kw, kw, kw,kw, kw, kw, kw, kw, kw, kw, kw, kw).Count(&count).Error
//	}else {
//		err = db.DbLink.Debug().Raw(sql,pageSize,offset).Scan(&strainIds).Error
//		if err != nil {
//			log.Error(err)
//			return nil,0
//		}
//		err = db.DbLink.Debug().Raw(countSql).Count(&count).Error
//	}
//	if err != nil {
//		log.Error(err)
//		return nil,0
//	}
//	return strainIds,count
//}
