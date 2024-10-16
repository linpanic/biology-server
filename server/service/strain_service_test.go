package service

import (
	"encoding/json"
	"fmt"
	"github.com/linpanic/biology-server/caches"
	"github.com/linpanic/biology-server/db"
	"github.com/linpanic/biology-server/dto"
	"github.com/linpanic/biology-server/logs"
	log "github.com/sirupsen/logrus"
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

//func InitNumber() {
//	number := dao.GetMaxStrainNumber()
//	if number == "" {
//		caches.InitNumber(0)
//		return
//	}
//	number = strings.TrimLeft(number, "#")
//	formatInt, err := strconv.ParseInt(number, 10, 64)
//	if err != nil {
//		log.Error(err)
//		panic(err)
//	}
//	caches.InitNumber(formatInt)
//}

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
	NewSerial := func(i int) string {
		return fmt.Sprintf(SERIAL, i)
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

	NewStrinAnnotate := func(i int) []string {
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
		strainAnnotate := NewStrinAnnotate(i)
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
			Name:     alleleName,
			Annotate: alleleAnnotate,
			Extra:    alleleExtra,
			Genome:   genomeName,
			Serial:   serial,
		})
		reqs = append(reqs, req)
	}

	for i := 100; i < 200; i++ {
		var req dto.StrainAddReq
		strainName := NewStrain(i)
		number := caches.GetNumber()
		sn := NewShortName(i)
		strainAnnotate := NewStrinAnnotate(i)
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
				Name:     alleleName,
				Annotate: alleleAnnotate,
				Extra:    alleleExtra,
				Genome:   genomeName,
				Serial:   serial,
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

//
//func TestAddByRecord(t *testing.T) {
//	excelFileName := "./hmu.xlsx"
//	xlFile, err := xlsx.OpenFile(excelFileName)
//	if err != nil {
//		log.Errorf("Error opening Excel file: %s", err)
//		return
//	}
//
//	var reqs []dto.StrainAddReq
//
//	for _, v := range xlFile.Sheets {
//		for i2, row := range v.Rows {
//			if i2 == 0 {
//				continue
//			}
//			var req dto.StrainAddReq
//			req.Number = row.Cells[0].String()
//
//			snStr := row.Cells[1].String()
//			snStr = strings.TrimSpace(snStr)
//			if snStr!= "" {
//				req.ShortName = append(req.ShortName, snStr)
//			}
//			//
//			//snStr = strings.ReplaceAll(snStr,"；",";")
//			//
//			//snSp := strings.Split(snStr, ";")
//			//for _, n := range snSp {
//			//	n = strings.TrimSpace(n)
//			//	if n == "" {
//			//		continue
//			//	}
//			//}
//
//			req.StrainName = row.Cells[2].String()
//			if len(row.Cells) > 5 {
//				for i3 := 6; i3 < len(row.Cells); i3++ {
//					req.StrainAnnotate = append(req.StrainAnnotate, row.Cells[i3].String())
//				}
//			}
//
//			alNameStr := row.Cells[3].String()
//			geStr := row.Cells[4].String()
//
//			alNameStr = strings.ReplaceAll(alNameStr, "；", ";")
//			alNameStr = strings.ReplaceAll(alNameStr, "：", ";")
//			alNameStr = strings.ReplaceAll(alNameStr, ":", ";")
//			alNameStr = strings.ReplaceAll(alNameStr, ";;", ";")
//			alNameStr = strings.ReplaceAll(alNameStr, ";;", ";")
//			alNameStr = strings.ReplaceAll(alNameStr, ";;", ";")
//			alNameStr = strings.ReplaceAll(alNameStr, ";;", ";")
//			geStr = strings.ReplaceAll(geStr, "；", ";")
//			geStr = strings.ReplaceAll(geStr, "{", "[")
//			geStr = strings.ReplaceAll(geStr, "}", "]")
//
//			var alArr, geArr []string
//
//			if alNameStr != "" {
//				sp := strings.Split(alNameStr, ";")
//				for _, name := range sp {
//					name = strings.TrimSpace(name)
//					if name == "" {
//						continue
//					}
//					alArr = append(alArr, name)
//				}
//			}
//
//			if geStr != "" {
//				geStr = strings.TrimSpace(geStr)
//				rs := []rune(geStr)
//				isIn := false
//				st := 0
//				for rIndex, r := range rs {
//					c := string(r)
//					if c == "" {
//						continue
//					}
//					if c == "[" {
//						isIn = true
//					}
//					if c == "]" {
//						isIn = false
//					}
//					if c == ";" && !isIn {
//						if string(rs[rIndex-1]) == ";" {
//							st = rIndex+1
//							continue
//						}
//						geArr = append(geArr, string(rs[st:rIndex]))
//						st = rIndex + 1
//					}
//					if len(rs)-1 == rIndex {
//						geArr = append(geArr, string(rs[st:rIndex+1]))
//					}
//				}
//			}
//
//			if len(alArr) == 0 && len(geArr) > 0 {
//				for ai := 0; ai < len(geArr); ai++ {
//					alArr = append(alArr, "")
//				}
//			}
//
//			if len(alArr) > 0 && len(geArr) == 0 {
//				for ai := 0; ai < len(alArr); ai++ {
//					geArr = append(geArr, "")
//				}
//			}
//
//			if len(alArr) != len(geArr) {
//				log.Error(req.Number)
//				panic("arr length is error")
//			}
//
//			if len(alArr) > 0 || len(geArr) > 0 {
//				for idx := 0; idx < len(alArr); idx++ {
//					var allele dto.Allele
//					allele.Name = alArr[idx]
//					allele.Genome = geArr[idx]
//					req.Allele = append(req.Allele, allele)
//				}
//			}
//
//			if row.Cells[5] != nil {
//				req.StrainExtra = append(req.StrainExtra, dto.ExtraInfo{
//					ExtraKey: "存在状态",
//					ExtraVal: row.Cells[5].String(),
//				})
//			}
//
//			//if len(row.Cells) >6 {
//			//	allele.Extra = append(allele.Extra,dto.ExtraInfo{
//			//		ExtraKey: "列1",
//			//		ExtraVal:  row.Cells[6].String(),
//			//	})
//			//}else {
//			//	allele.Extra = append(allele.Extra,dto.ExtraInfo{
//			//		ExtraKey: "列1",
//			//		ExtraVal:  "",
//			//	})
//			//}
//
//			//if len(row.Cells) >5 {
//			//	allele.Annotate = []string{row.Cells[5].String()}
//			//}
//			reqs = append(reqs, req)
//		}
//
//		var s StrainService
//		for _, r := range reqs {
//			result := s.Add(r, 1)
//			log.Info(result.Message)
//		}
//		break
//	}
//
//}

func TestGet(t *testing.T) {

	var s StrainService
	list := s.List(dto.StrainListReq{
		PageReq: dto.PageReq{
			PageNo:   1,
			PageSize: 10,
		},
		Field: "number",
		Order: "desc",
	})

	marshal, _ := json.Marshal(list)
	log.Info(string(marshal))
}
