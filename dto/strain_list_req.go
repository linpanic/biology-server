package dto

import (
	"github.com/linpanic/biology-server/caches"
	"github.com/linpanic/biology-server/etype"
)

type StrainListReq struct {
	PageReq        //分页
	Key     string `json:"key,omitempty"`   //搜索的关键词
	Field   string `json:"field,omitempty"` //字段
	Order   string `json:"order,omitempty"` //排序方式
}

func (s *StrainListReq) Verify() error {
	if s.PageNo < 0 || s.PageSize < 0 {
		return etype.PageError
	}

	if s.Order != "" {
		if s.Order != "desc" && s.Order != "asc" {
			return etype.OrderError
		}
	} else {
		s.Order = "desc"
	}

	if s.Field != "" {
		ok := false
		for _, v := range caches.StrainAlleleField {
			if v == s.Field {
				ok = true
				break
			}
		}

		if !ok {
			return etype.FieldError
		}
	}

	return nil
}
