package utils

import (
	"reflect"
)

func FieldEmpty(data any) bool {
	value := reflect.ValueOf(data)
	switch value.Kind() {
	case reflect.Pointer:
		value = value.Elem()
	}
	for i := 0; i < value.NumField(); i++ {
		if value.Field(i).IsZero() {
			return false
		}
	}
	return true
}
