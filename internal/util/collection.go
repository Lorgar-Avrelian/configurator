package util

import (
	"reflect"
	"sort"
)

func SortSliceByID[T any](slice []T) {
	var sliceLen int
	var firstElem reflect.Value
	var fieldName string
	var hasField bool
	sliceLen = len(slice)
	if sliceLen == 0 {
		return
	}
	firstElem = reflect.ValueOf(slice[0])
	if firstElem.Kind() == reflect.Ptr {
		firstElem = firstElem.Elem()
	}
	fieldName = "ID"
	hasField = firstElem.FieldByName(fieldName).IsValid()
	if !hasField {
		fieldName = "Id"
		hasField = firstElem.FieldByName(fieldName).IsValid()
	}
	if !hasField {
		return
	}
	sort.Slice(slice, func(i, j int) bool {
		var vI reflect.Value
		var vJ reflect.Value
		var idI int64
		var idJ int64
		var isGreater bool
		vI = reflect.ValueOf(slice[i])
		vJ = reflect.ValueOf(slice[j])
		if vI.Kind() == reflect.Ptr {
			vI = vI.Elem()
		}
		if vJ.Kind() == reflect.Ptr {
			vJ = vJ.Elem()
		}
		idI = vI.FieldByName(fieldName).Int()
		idJ = vJ.FieldByName(fieldName).Int()
		isGreater = idI > idJ
		return !isGreater
	})
}
