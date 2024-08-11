package service

import (
	"github.com/linpanic/biology-server/caches"
	"github.com/linpanic/biology-server/cst"
	"github.com/linpanic/biology-server/dao"
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/dto"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	StrainServiceInstance StrainService
	l                     = new(sync.Mutex)
)

type StrainService struct{}

// 获取序列号
func (s *StrainService) GetNumber(req dto.StrainNumberReq) dto.Result {
	l.Unlock()
	defer l.Unlock()
	err := req.Verify()
	if err != nil {
		log.Error(err)
		return dto.NewErrResult(cst.VERIFY_ERROR, err.Error())
	}
	return dto.NewOKResult(dto.StrainNumberResp{caches.GetNumber()})
}

// 新增品系
func (s *StrainService) Add(req dto.StrainAddReq, userId int64) dto.Result {
	l.Lock()
	defer func() {
		InitNumber()
		l.Unlock()
	}()

	err := req.Verify()
	if err != nil {
		log.Error(err)
		return dto.NewErrResult(cst.VERIFY_ERROR, err.Error())
	}

	if req.Number == "" {
		return dto.NewErrResult(cst.NUMBER_NULL, "序列号不存在")
	}

	if req.StrainName != "" {
		//校验品系名是否重复
		strain := dao.SelectStrainByName(req.StrainName)
		if strain != nil {
			return dto.NewErrResult(cst.STRAIN_EXIST, "品系名已存在")
		}
	}

	//校验序列号是否重复
	strain := dao.SelectStrainByNum(req.Number)
	if strain != nil {
		return dto.NewErrResult(cst.NUMBER_EXIST, "序列号已存在")
	}

	now := time.Now().Unix()

	//新增品系
	tx := db.DbLink.Begin()
	strain, err = dao.CreateStrain(tx, req.Number, req.StrainName, userId, now)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return dto.NewErrResult(cst.DAO_ERROR, err.Error())
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
			return dto.NewErrResult(cst.DAO_ERROR, err.Error())
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
			return dto.NewErrResult(cst.DAO_ERROR, err.Error())
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
			return dto.NewErrResult(cst.DAO_ERROR, err.Error())
		}
	}

	//处理基因
	if len(req.Allele) > 0 {

		for _, v := range req.Allele {
			//新增注解

			allele, err := dao.CreateAllele(tx, strain.Id, v, userId, now)
			if err != nil {
				log.Error(err)
				tx.Rollback()
				return dto.NewErrResult(cst.DAO_ERROR, err.Error())
			}
			if len(v.Annotate) > 0 {
				for i, v2 := range v.Annotate {
					v2 = strings.TrimSpace(v2)
					if v2 != "" {
						v.Annotate[i] = v2
					}
				}
				err = dao.CreateAlleleAnnotate(tx, allele.Id, v.Annotate, userId, now)
				if err != nil {
					log.Error(err)
					tx.Rollback()
					return dto.NewErrResult(cst.DAO_ERROR, err.Error())
				}
			}

			//新增额外信息
			if len(v.Extra) > 0 {
				for i, v2 := range v.Extra {
					v2.ExtraKey = strings.TrimSpace(v2.ExtraKey)
					v2.ExtraVal = strings.TrimSpace(v2.ExtraVal)
					if v2.ExtraKey != "" || v2.ExtraVal != "" {
						v.Extra[i] = v2
					}
				}
				err = dao.CreateAlleleExtra(tx, allele.Id, v.Extra, userId, now)
				if err != nil {
					log.Error(err)
					tx.Rollback()
					return dto.NewErrResult(cst.DAO_ERROR, err.Error())
				}
			}

		}
	}

	tx.Commit()
	return dto.NewOKResult(nil)
}

// 展示列表
func (s *StrainService) List(req dto.StrainListReq) dto.Result {
	err := req.Verify()
	if err != nil {
		log.Error(err)
		return dto.NewErrResult(cst.VERIFY_ERROR, err.Error())
	}

	strainAlleles, count := dao.SelectStrainAndAllele(req.Key, req.Field, req.Order, req.PageNo, req.PageSize)
	if len(strainAlleles) == 0 || count == 0 {
		return dto.NewOKResult([]struct{}{})
	}

	var resp dto.StrainListResp

	var list []dto.Strain
	for _, v := range strainAlleles {
		var strain dto.Strain
		strain.Id = v.Id
		strain.StrainName = v.StrainName
		strain.Number = v.Number

		if v.ShortName != "" {
			split := strings.Split(v.ShortName, "△")
			strain.ShortName = split
		}

		if v.StrainAnnotate != "" {
			split := strings.Split(v.StrainAnnotate, "△")
			strain.StrainAnnotate = split
		}

		if v.StrainExtraKey != "" {
			split := strings.Split(v.StrainExtraKey, "△")
			for _, v2 := range split {
				strain.StrainExtra = append(strain.StrainExtra, dto.ExtraInfo{
					ExtraKey: v2,
				})
			}
		}

		if v.StrainExtraValue != "" {
			split := strings.Split(v.StrainExtraValue, "△")
			for i, v2 := range split {
				strain.StrainExtra[i].ExtraVal = v2
			}
		}

		if v.AlleleId != "" {
			aIds := strings.Split(v.AlleleId, "△")
			aNames := strings.Split(v.AlleleName, "△")
			genome := strings.Split(v.Genome, "△")
			serials := strings.Split(v.Serial, "△")

			aExtraKeys := strings.Split(v.AlleleExtraKey, "△")
			aExtraVals := strings.Split(v.AlleleExtraValue, "△")
			aAnnotates := strings.Split(v.AAnnotate, "△")

			if len(aIds) == 0 {
				continue
			}
			alesMap := make(map[int64]dto.Allele)

			for i, v2 := range aIds {
				if v2 == "" {
					continue
				}
				var ale dto.Allele
				aId, _ := strconv.ParseInt(v2, 10, 64)
				ale.Id = aId
				ale.Name = aNames[i]
				ale.Genome = genome[i]
				ale.Serial = serials[i]
				alesMap[aId] = ale
			}
			if v.AlleleExtraKey != "" && len(aExtraKeys) > 0 {
				for i := 0; i < len(aExtraKeys); i++ {
					idStr, key, _ := strings.Cut(aExtraKeys[i], "☆")
					_, value, _ := strings.Cut(aExtraVals[i], "☆")
					aId, _ := strconv.ParseInt(idStr, 10, 64)
					ale := alesMap[aId]

					ale.Extra = append(ale.Extra, dto.ExtraInfo{
						ExtraKey: key,
						ExtraVal: value,
					})
					alesMap[aId] = ale
				}

				//for _, key := range aExtraKeys {
				//	if key == "" {
				//		continue
				//	}
				//	before, after, _ := strings.Cut(key, "☆")
				//	aId, _ := strconv.ParseInt(before, 10, 64)
				//	ale := alesMap[aId]
				//
				//	ale.Extra = append(ale.Extra, dto.ExtraInfo{
				//		ExtraKey: after,
				//	})
				//	alesMap[aId] = ale
				//}
			}

			//if len(aExtraVals) > 0 {
			//	for i, v3 := range aExtraVals {
			//		if v3 == "" {
			//			continue
			//		}
			//		before, after, _ := strings.Cut(v3, "☆")
			//		aId, _ := strconv.ParseInt(before, 10, 64)
			//		ale := alesMap[aId]
			//		ale.Extra[i].ExtraVal = after
			//		alesMap[aId] = ale
			//	}
			//}

			if v.AAnnotate != "" && len(aAnnotates) > 0 {
				for _, v3 := range aAnnotates {
					if v3 == "" {
						continue
					}
					before, after, _ := strings.Cut(v3, "☆")
					aId, _ := strconv.ParseInt(before, 10, 64)
					ale := alesMap[aId]
					ale.Annotate = append(ale.Annotate, after)
					alesMap[aId] = ale
				}
			}

			for _, v2 := range alesMap {
				strain.Allele = append(strain.Allele, v2)
			}
		}
		list = append(list, strain)
	}
	resp.PageNo = req.PageNo
	resp.PageSize = req.PageSize
	resp.Total = count
	resp.StrainList = list
	return dto.NewOKResult(resp)
}

// 修改品系数据
func (s *StrainService) Update(req dto.StrainUpdateReq, userId int64) dto.Result {
	strain := dao.SelectOneStrain(req.Id, 0)
	if strain == nil {
		return dto.NewErrResult(cst.DAO_ERROR, "找不到该品系ID")
	}

	tx := db.DbLink.Begin()
	err := dao.DeleteShortName(tx, req.Id)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return dto.NewErrResult(cst.DAO_ERROR, err.Error())
	}

	err = dao.DeleteStrainAnnotate(tx, req.Id)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return dto.NewErrResult(cst.DAO_ERROR, err.Error())
	}

	err = dao.DeleteStrainExtra(tx, req.Id)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return dto.NewErrResult(cst.DAO_ERROR, err.Error())
	}

	now := time.Now().Unix()
	err = dao.UpdateStrain(tx, req.Id, req.StrainName, now)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return dto.NewErrResult(cst.DAO_ERROR, err.Error())
	}

	err = dao.CreateStrainExtra(tx, req.Id, req.StrainExtra, userId, now)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return dto.NewErrResult(cst.DAO_ERROR, err.Error())
	}

	err = dao.CreateShortName(tx, req.Id, req.ShortName, userId, now)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return dto.NewErrResult(cst.DAO_ERROR, err.Error())
	}

	err = dao.CreateStrainAnnotate(tx, req.Id, req.StrainAnnotate, userId, now)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return dto.NewErrResult(cst.DAO_ERROR, err.Error())
	}

	var alleleReq dto.AlleleUpdateReq
	alleleReq.Id = req.Id
	alleleReq.Allele = req.Allele

	err = AlleleServiceInstance.UpdateWithStrain(tx, alleleReq, userId)
	if err != nil {
		tx.Rollback()
		return dto.NewErrResult(cst.DAO_ERROR, err.Error())
	}

	tx.Commit()
	return dto.NewOKResult(nil)
}

func (s *StrainService) Delete(req dto.StrainDelReq) dto.Result {
	err := req.Verify()
	if err != nil {
		log.Error(err)
		return dto.NewErrResult(cst.VERIFY_ERROR, err.Error())
	}
	tx := db.DbLink.Begin()
	err = dao.DeleteOneStrain(tx, req.StrainId)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return dto.NewErrResult(cst.DAO_ERROR, err.Error())
	}
	tx.Commit()
	return dto.NewOKResult(nil)
}
