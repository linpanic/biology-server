package service

import (
	"github.com/linpanic/biology-server/dao"
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/dto"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

var (
	AlleleServiceInstance AlleleService
)

type AlleleService struct{}

func (a *AlleleService) AlleleSearch(req dto.AlleleSearchReq) dto.Result {
	err := req.Verify()
	if err != nil {
		return dto.NewErrResult(err.Error())
	}
	alleles := dao.SelectAlleleByName(req.Name)
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

func (a *AlleleService) Update(req dto.AlleleUpdateReq, userId int64) dto.Result {

	tx := db.DbLink.Begin()

	var ids []int64

	for _, v := range req.Allele {
		ids = append(ids, v.Id)
	}

	now := time.Now().Unix()

	err := dao.DeleteAlleleAnnotates(tx, ids)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return dto.NewErrResult(err.Error())
	}

	err = dao.DeleteAlleleExtras(tx, ids)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return dto.NewErrResult(err.Error())
	}

	var creates []dto.Allele
	for _, v := range req.Allele {
		if v.Id != 0 {
			err = dao.UpdateAllele(tx, v.Id, v.Name, v.Genome, v.Serial, now)
			if err != nil {
				log.Error(err)
				tx.Rollback()
				return dto.NewErrResult(err.Error())
			}
		} else {
			creates = append(creates, v)
		}
	}
	if len(creates) > 0 {
		result, err := dao.CreateAlleles(tx, req.Id, creates, userId, now)
		if err != nil {
			log.Error(err)
			tx.Rollback()
			return dto.NewErrResult(err.Error())
		}
		for i, v := range result {
			creates[i].Id = v.Id
		}
	}

	req.Allele = append(req.Allele, creates...)
	err = dao.CreateAllelesAnnotate(tx, req.Allele, userId, now)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return dto.NewErrResult(err.Error())
	}

	err = dao.CreateAllelesExtra(tx, req.Allele, userId, now)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return dto.NewErrResult(err.Error())
	}
	tx.Commit()
	return dto.NewOKResult(nil)
}

func (a *AlleleService) UpdateWithStrain(tx *gorm.DB, req dto.AlleleUpdateReq, userId int64) error {

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
