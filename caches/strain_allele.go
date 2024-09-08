package caches

import (
	"github.com/linpanic/biology-server/model"
	"reflect"
	"strings"
	"sync"
)

var (
	StrainAlleleField []string
	AlleleField       []string
	initOnce          sync.Once
)

func InitStrainAlleleField() {
	initOnce.Do(func() {
		var s model.StrainAllele
		r := reflect.TypeOf(s)
		for i := 0; i < r.NumField(); i++ {
			tag := r.Field(i).Tag
			tn := tag.Get("json")
			before, _, found := strings.Cut(tn, ",")
			if found {
				StrainAlleleField = append(StrainAlleleField, before)
			}
		}

		var a model.AlleleAll
		r = reflect.TypeOf(a)
		for i := 0; i < r.NumField(); i++ {
			tag := r.Field(i).Tag
			tn := tag.Get("json")
			before, _, found := strings.Cut(tn, ",")
			if found {
				AlleleField = append(AlleleField, before)
			}
		}
	})
}
