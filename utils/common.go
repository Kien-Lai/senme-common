package utils

import (
	"reflect"
	"strings"
)

const (
	Empty           = ""
	Space           = " "
	DefaultLanguage = "vi"
)

type DefaultRequestQuery struct {
	Language string `form:"language, default=vi"`
	Page     int    `form:"page, default=1"`
	PageSize int    `form:"pageSize, default=20"`
}

func IsBlank(s string) bool {

	return len(strings.Trim(s, Space)) == 0
}

func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

func IsNil(s interface{}) bool {
	if pointer := &s; pointer == nil {
		return true
	}
	return false
}

func IsNotNil(s interface{}) bool {
	return !IsNil(s)
}

func Contains(s []string, e string) bool {
	if IsNil(s) {
		return false
	}
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsDifference(s []string, e string) bool {
	for _, a := range s {
		if a != e {
			return true
		}
	}
	return false
}

func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}
