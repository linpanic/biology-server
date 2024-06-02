package service

import (
	"github.com/linpanic/biology-server/caches"
	"github.com/linpanic/biology-server/dto"
	log "github.com/sirupsen/logrus"
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
