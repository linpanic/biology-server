package service

import (
	"github.com/linpanic/biology-server/caches"
	"github.com/linpanic/biology-server/dao"
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/dto"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type StrainService struct{}

// 获取序列号
func (s *StrainService) GetNumber(req dto.StrainNumberReq) dto.Result {
	err := req.Verify()
	if err != nil {
		log.Error(err)
		return dto.NewErrResult(err.Error())
	}
	return dto.NewOKResult(dto.StrainNumberResp{caches.GetNumber()})
}

// 新增品系
func (s *StrainService) Add(req dto.StrainAddReq, userId int64) dto.Result {
	err := req.Verify()
	if err != nil {
		log.Error(err)
		return dto.NewErrResult(err.Error())
	}

	if req.Number == "" {
		return dto.NewErrResult("序列号不存在")
	}

	if req.StrainName != "" {
		//校验品系名是否重复
		strain := dao.SelectStrainByName(req.StrainName)
		if strain != nil {
			return dto.NewErrResult("品系名已存在")
		}
	}

	//校验序列号是否重复
	strain := dao.SelectStrainByNum(req.Number)
	if strain != nil {
		return dto.NewErrResult("序列号已存在")
	}

	now := time.Now().Unix()

	//新增品系
	tx := db.DbLink.Begin()
	strain, err = dao.CreateStrain(tx, req.Number, req.StrainName, userId, now)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return dto.NewErrResult(err.Error())
	}

	//新增简称
	if len(req.ShortName) > 0 {
		for i, v := range req.ShortName {
			v = strings.TrimSpace(v)
			if v != "" {
				req.ShortName[i] = v
			}
		}
		err = dao.CreateShortName(tx, strain.Id, req.ShortName, userId, now)
		if err != nil {
			log.Error(err)
			tx.Rollback()
			return dto.NewErrResult(err.Error())
		}
	}

	//新增注解
	if len(req.StrainAnnotate) > 0 {
		for i, v := range req.StrainAnnotate {
			v = strings.TrimSpace(v)
			if v != "" {
				req.StrainAnnotate[i] = v
			}
		}
		err = dao.CreateStrainAnnotate(tx, strain.Id, req.StrainAnnotate, userId, now)
		if err != nil {
			log.Error(err)
			tx.Rollback()
			return dto.NewErrResult(err.Error())
		}
	}

	//新增额外信息
	if len(req.StrainExtra) > 0 {
		for i, v := range req.StrainExtra {
			v.ExtraKey = strings.TrimSpace(v.ExtraKey)
			v.ExtraVal = strings.TrimSpace(v.ExtraVal)
			if v.ExtraKey != "" || v.ExtraVal != "" {
				req.StrainExtra[i] = v
			}
		}
		err = dao.CreateStrainExtra(tx, strain.Id, req.StrainExtra, userId, now)
		if err != nil {
			log.Error(err)
			tx.Rollback()
			return dto.NewErrResult(err.Error())
		}
	}

	//处理基因
	if len(req.Allele) > 0 {
		for _, v := range req.Allele {
			v.AlleleName = strings.TrimSpace(v.AlleleName)
			if v.AlleleName != "" {

				//新增基因
				allele, err := dao.CreateAlleleName(tx, strain.Id, v.AlleleName, userId, now)
				if err != nil {
					log.Error(err)
					tx.Rollback()
					return dto.NewErrResult(err.Error())
				}

				//新增注解
				if len(v.AlleleNameAnnotate) > 0 {
					for i, v2 := range v.AlleleNameAnnotate {
						v2 = strings.TrimSpace(v2)
						if v2 != "" {
							v.AlleleNameAnnotate[i] = v2
						}
					}
					err = dao.CreateAlleleNameAnnotate(tx, allele.Id, v.AlleleNameAnnotate, userId, now)
					if err != nil {
						log.Error(err)
						tx.Rollback()
						return dto.NewErrResult(err.Error())
					}
				}

				//新增额外信息
				if len(v.AlleleNameExtra) > 0 {
					for i, v2 := range req.StrainExtra {
						v2.ExtraKey = strings.TrimSpace(v2.ExtraKey)
						v2.ExtraVal = strings.TrimSpace(v2.ExtraVal)
						if v2.ExtraKey != "" || v2.ExtraVal != "" {
							v.AlleleNameExtra[i] = v2
						}
					}
					err = dao.CreateAlleleNameExtra(tx, allele.Id, v.AlleleNameExtra, userId, now)
					if err != nil {
						log.Error(err)
						tx.Rollback()
						return dto.NewErrResult(err.Error())
					}
				}

				//基因修饰情况
				v.GenomeName = strings.TrimSpace(v.GenomeName)
				if v.GenomeName != "" {
					genome, err := dao.CreateGenome(tx, allele.Id, v.GenomeName, userId, now)
					if err != nil {
						log.Error(err)
						tx.Rollback()
						return dto.NewErrResult(err.Error())
					}

					//染色体信息
					if len(v.Serial) > 0 {
						var serials []string
						for _, v2 := range v.Serial {
							v2.Serial = strings.TrimSpace(v2.Serial)
							if v2.Serial != "" {
								serials = append(serials, v2.Serial)
							}
						}
						err = dao.CreateChromsome(tx, genome.Id, serials, userId, now)
						if err != nil {
							log.Error(err)
							tx.Rollback()
							return dto.NewErrResult(err.Error())
						}
					}
				}
			}
		}
	}

	tx.Commit()
	return dto.NewOKResult(nil)
}

////展示列表
//func (s *StrainService) List(req dto.StrainListReq) dto.Result {
//	err := req.Verify()
//	if err != nil {
//		log.Error(err)
//		return dto.NewErrResult(err.Error())
//	}
//
//
//
//}
