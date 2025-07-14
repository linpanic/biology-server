package service

import (
	"github.com/linpanic/biology-server/cst"
	"github.com/linpanic/biology-server/dao"
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/dto"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
	"time"
)

var (
	AlleleServiceInstance AlleleService
)

type AlleleService struct {
	db *gorm.DB
}

func NewAlleleService(db *gorm.DB) *AlleleService {
	return &AlleleService{
		db: db,
	}
}

// 增加基因
func (a *AlleleService) Add(req dto.AlleleAddReq, userId int64) dto.Result {
	//删除注解和额外信息
	tx := db.DbLink.Begin()

	now := time.Now().Unix()

	ae := dto.Allele{
		Name:   req.Name,
		Genome: req.Genome,
		Serial: req.Serial,
	}

	allele, err := dao.CreateAllele(tx, 0, ae, userId, now)
	if err != nil {
		log.Error(err)
		if err != nil {
			log.Error(err)
			tx.Rollback()
			return dto.NewErrResult(cst.DAO_ERROR, err.Error())
		}
	}

	//添加基因额外信息和注解
	if len(req.Extra) > 0 {
		err = dao.CreateAlleleExtra(tx, allele.Id, req.Extra, userId, now)
		if err != nil {
			log.Error(err)
			tx.Rollback()
			return dto.NewErrResult(cst.DAO_ERROR, err.Error())
		}
	}

	if len(req.Annotate) > 0 {
		err = dao.CreateAlleleAnnotate(tx, allele.Id, req.Annotate, userId, now)
		if err != nil {
			log.Error(err)
			tx.Rollback()
			return dto.NewErrResult(cst.DAO_ERROR, err.Error())
		}
	}

	tx.Commit()
	return dto.NewOKResult(nil)
}

// 普通搜索
func (a *AlleleService) AlleleSearch(req dto.AlleleSearchReq) dto.Result {
	err := req.Verify()
	if err != nil {
		return dto.NewErrResult(cst.VERIFY_ERROR, err.Error())
	}
	alleles := dao.SelectAlleleByName(a.db, req.Name)
	var resp dto.AlleleSearchResp
	for _, v := range alleles {
		resp.Allele = append(resp.Allele, dto.Allele{
			Id:     v.Id,
			Name:   v.Name,
			Genome: v.Genome,
			Serial: v.Serial,
		})
	}
	return dto.NewOKResult(resp)
}

// 全字段搜索
func (a *AlleleService) AlleleAllSearch(req dto.AlleleListReq) dto.Result {
	err := req.Verify()
	if err != nil {
		return dto.NewErrResult(cst.VERIFY_ERROR, err.Error())
	}
	alleles, count := dao.SelectAlleleByAll(a.db, req.Key, req.Field, req.Order, req.PageNo, req.PageSize)
	var resp dto.AlleleAllListResp
	resp.PageNo = req.PageNo
	resp.PageSize = req.PageSize
	resp.Total = count
	var list []dto.AlleleAll

	for _, v := range alleles {
		var alleleAll dto.AlleleAll
		alleleAll.Id = v.AlleleId
		alleleAll.Name = v.Name
		alleleAll.Genome = v.Genome
		alleleAll.Serial = v.Serial

		extraKeys := strings.Split(v.ExtraKey, "△")
		extraVals := strings.Split(v.ExtraVal, "△")
		annotates := strings.Split(v.Annotate, "△")

		if v.ExtraKey != "" && len(extraKeys) > 0 {
			for i := 0; i < len(extraKeys); i++ {
				eKey := extraKeys[i]
				eVal := extraVals[i]
				alleleAll.Extra = append(alleleAll.Extra, dto.ExtraInfo{
					ExtraKey: eKey,
					ExtraVal: eVal,
				})
			}
		}

		if v.Annotate != "" && len(annotates) > 0 {
			alleleAll.Annotate = annotates
		}
		list = append(list, alleleAll)
	}
	resp.Allele = list
	return dto.NewOKResult(resp)
}

// 更新单独
func (a *AlleleService) Update(req dto.AlleleUpdateReq, userId int64) dto.Result {
	//校验数据是否存在
	allele := dao.SelectAlleleById(a.db, req.Id)
	if allele == nil {
		log.Error("找不到对应ID基因:", req.Id)
		return dto.NewErrResult(cst.DAO_ERROR, "该ID不存在")
	}

	//删除注解和额外信息
	tx := db.DbLink.Begin()

	ids := []int64{req.Id}

	now := time.Now().Unix()

	err := dao.DeleteAlleleAnnotates(tx, ids)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return dto.NewErrResult(cst.DAO_ERROR, err.Error())
	}

	err = dao.DeleteAlleleExtras(tx, ids)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return dto.NewErrResult(cst.DAO_ERROR, err.Error())
	}

	//修改基因信息
	err = dao.UpdateAllele(tx, req.Id, req.Name, req.Genome, req.Serial, now)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return dto.NewErrResult(cst.DAO_ERROR, err.Error())
	}

	//添加基因额外信息和注解
	if len(req.Extra) > 0 {
		err = dao.CreateAlleleExtra(tx, req.Id, req.Extra, userId, now)
		if err != nil {
			log.Error(err)
			tx.Rollback()
			return dto.NewErrResult(cst.DAO_ERROR, err.Error())
		}
	}

	if len(req.Annotate) > 0 {
		err = dao.CreateAlleleAnnotate(tx, req.Id, req.Annotate, userId, now)
		if err != nil {
			log.Error(err)
			tx.Rollback()
			return dto.NewErrResult(cst.DAO_ERROR, err.Error())
		}
	}

	tx.Commit()
	return dto.NewOKResult(nil)
}

// 删除基因
func (a *AlleleService) Delete(req dto.AlleleDelReq, userId int64) dto.Result {
	err := dao.DeleteStrain(a.db, []int64{req.Id})
	if err != nil {
		if err != nil {
			log.Error(err)
			return dto.NewErrResult(cst.DAO_ERROR, err.Error())
		}
	}
	return dto.NewOKResult(nil)
}

// 通过品系更新
func (a *AlleleService) UpdateWithStrain(tx *gorm.DB, req dto.StrainAlleleUpdateReq, userId int64) error {

	var ids []int64

	for _, v := range req.Allele {
		if v.Id == 0 {
			continue
		}
		ids = append(ids, v.Id)
	}

	now := time.Now().Unix()

	err := dao.DeleteAlleleAnnotates(tx, ids)
	if err != nil {
		log.Error(err)
		//tx.Rollback()
		return err
	}

	err = dao.DeleteAlleleExtras(tx, ids)
	if err != nil {
		log.Error(err)
		//tx.Rollback()
		return err
	}

	var creates []dto.Allele
	for _, v := range req.Allele {
		if v.Id != 0 {
			err = dao.UpdateAllele(tx, v.Id, v.Name, v.Genome, v.Serial, now)
			if err != nil {
				log.Error(err)
				//tx.Rollback()
				return err
			}
		} else {
			creates = append(creates, v)
		}
	}

	if len(creates) > 0 {
		result, err := dao.CreateAlleles(tx, req.Id, creates, userId, now)
		if err != nil {
			log.Error(err)
			//tx.Rollback()
			return err
		}
		for i, v := range result {
			creates[i].Id = v.Id
		}
	}

	req.Allele = append(req.Allele, creates...)
	err = dao.CreateAllelesAnnotate(tx, req.Allele, userId, now)
	if err != nil {
		log.Error(err)
		//tx.Rollback()
		return err
	}

	err = dao.CreateAllelesExtra(tx, req.Allele, userId, now)
	if err != nil {
		log.Error(err)
		//tx.Rollback()
		return err
	}
	return nil
}
