package service

import (
	"fmt"
	"github.com/linpanic/biology-server/caches"
	"github.com/linpanic/biology-server/dao"
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/dto"
	"github.com/linpanic/biology-server/logs"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"testing"
)

func init() {
	//初始化日志
	logs.LogInit()

	//初始化数据库链接
	db.DbInit()

	//初始化序列
	InitNumber()
}

func InitNumber() {
	number := dao.GetMaxStrainNumber()
	if number == "" {
		caches.InitNumber(0)
		return
	}
	number = strings.TrimLeft(number, "#")
	formatInt, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	caches.InitNumber(formatInt)
}

func TestAdd(t *testing.T) {
	STRAIN := "品系_%d"

	SHORT_NAME1 := "简称Version_%d"
	SHORT_NAME2 := "简称EX_%d"
	STRAIN_ANNOTATE_1 := "这是华为_%d"
	STRAIN_ANNOTATE_2 := "这是小米_%d"

	ALLELE_ANNOTATE_1 := "这是OPPO_%d"
	ALLELE_ANNOTATE_2 := "这是VIVO_%d"

	STRAIN_EXK := "品系作者"
	STRAIN_EXV := "品系终结者ID:%d"

	STRAIN_EXK2 := "品系数据来源"
	STRAIN_EXV2 := "品系第%d网百度来的"

	ALLELE_EXK := "基因作者"
	ALLELE_EXV := "基因终结者ID:%d"

	ALLELE_EXK2 := "基因数据来源"
	ALLELE_EXV2 := "基因第%d网百度来的"

	ALLELE_NAME := "基因名_%d"

	GENOME_NAME := "基因修饰情况_%d"

	NewGenome := func(i int) string {
		return fmt.Sprintf(GENOME_NAME, i)
	}

	SERIAL := "染色体_%d"
	NewSerial := func(i int) []dto.Serial {
		return []dto.Serial{
			{
				Serial: fmt.Sprintf(SERIAL, i),
			},
		}
	}

	NewAlleleName := func(i int) string {
		return fmt.Sprintf(ALLELE_NAME, i)
	}

	NewStrain := func(i int) string {
		return fmt.Sprintf(STRAIN, i)
	}

	NewShortName := func(i int) []string {
		switch i % 3 {
		case 0:
			return []string{fmt.Sprintf(SHORT_NAME1, i), fmt.Sprintf(SHORT_NAME2, i)}
		case 1:
			return nil
		default:
			return []string{fmt.Sprintf(SHORT_NAME1, i)}
		}
	}

	NewstrinAnnotate := func(i int) []string {
		switch i % 3 {
		case 0:
			return []string{fmt.Sprintf(STRAIN_ANNOTATE_1, i), fmt.Sprintf(STRAIN_ANNOTATE_2, i)}
		case 1:
			return nil
		default:
			return []string{fmt.Sprintf(STRAIN_ANNOTATE_1, i)}
		}
	}

	NewStainExtra := func(i int) []dto.ExtraInfo {
		switch i % 3 {
		case 0:
			return []dto.ExtraInfo{
				{
					ExtraKey: STRAIN_EXK,
					ExtraVal: fmt.Sprintf(STRAIN_EXV, i),
				},
				{
					ExtraKey: STRAIN_EXK2,
					ExtraVal: fmt.Sprintf(STRAIN_EXV2, i),
				}}
		case 1:
			return nil
		default:
			return []dto.ExtraInfo{
				{ExtraKey: STRAIN_EXK, ExtraVal: fmt.Sprintf(STRAIN_EXV, i)},
			}
		}
	}

	NewAlleleExtra := func(i int) []dto.ExtraInfo {
		switch i % 3 {
		case 0:
			return []dto.ExtraInfo{
				{
					ExtraKey: ALLELE_EXK,
					ExtraVal: fmt.Sprintf(ALLELE_EXV, i),
				},
				{
					ExtraKey: ALLELE_EXK2,
					ExtraVal: fmt.Sprintf(ALLELE_EXV2, i),
				}}
		case 1:
			return nil
		default:
			return []dto.ExtraInfo{
				{ExtraKey: ALLELE_EXK, ExtraVal: fmt.Sprintf(ALLELE_EXV, i)},
			}
		}
	}

	NewAlleleAnnotate := func(i int) []string {
		switch i % 3 {
		case 0:
			return []string{fmt.Sprintf(ALLELE_ANNOTATE_1, i), fmt.Sprintf(ALLELE_ANNOTATE_2, i)}
		case 1:
			return nil
		default:
			return []string{fmt.Sprintf(ALLELE_ANNOTATE_1, i)}
		}
	}

	var reqs []dto.StrainAddReq
	for i := 1; i < 100; i++ {
		var req dto.StrainAddReq
		strainName := NewStrain(i)
		number := caches.GetNumber()
		sn := NewShortName(i)
		strainAnnotate := NewstrinAnnotate(i)
		strainExtra := NewStainExtra(i)

		alleleName := NewAlleleName(i)
		alleleAnnotate := NewAlleleAnnotate(i)
		alleleExtra := NewAlleleExtra(i)
		genomeName := NewGenome(i)
		serial := NewSerial(i)

		req.StrainName = strainName
		req.Number = number
		req.ShortName = sn
		req.StrainAnnotate = strainAnnotate
		req.StrainExtra = strainExtra
		req.Allele = append(req.Allele, dto.Allele{
			AlleleName:         alleleName,
			AlleleNameAnnotate: alleleAnnotate,
			AlleleNameExtra:    alleleExtra,
			GenomeName:         genomeName,
			Serial:             serial,
		})
		reqs = append(reqs, req)
	}

	for i := 100; i < 200; i++ {
		var req dto.StrainAddReq
		strainName := NewStrain(i)
		number := caches.GetNumber()
		sn := NewShortName(i)
		strainAnnotate := NewstrinAnnotate(i)
		strainExtra := NewStainExtra(i)

		req.StrainName = strainName
		req.Number = number
		req.ShortName = sn
		req.StrainAnnotate = strainAnnotate
		req.StrainExtra = strainExtra

		for j := 0; j < 2; j++ {
			alleleName := NewAlleleName(i*100 + j)
			alleleAnnotate := NewAlleleAnnotate(i*100 + j)
			alleleExtra := NewAlleleExtra(i*100 + j)
			genomeName := NewGenome(i*100 + j)
			serial := NewSerial(i*100 + j)

			req.Allele = append(req.Allele, dto.Allele{
				AlleleName:         alleleName,
				AlleleNameAnnotate: alleleAnnotate,
				AlleleNameExtra:    alleleExtra,
				GenomeName:         genomeName,
				Serial:             serial,
			})
		}
		reqs = append(reqs, req)
	}

	var s StrainService
	for _, v := range reqs {
		result := s.Add(v, 1)
		log.Info(result.Message)
	}

}
