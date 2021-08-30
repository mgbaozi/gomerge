package gomerge

import (
	"fmt"
	"reflect"
)

func Merge(dst interface{}, src map[string]interface{}) (err error) {
	var processingField string
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s has wrong value type: %s", processingField, r)
		}
	}()
	dstValue := reflect.ValueOf(dst)
	dstType := dstValue.Elem().Type()
	for i := 0; i < dstType.NumField(); i++ {
		typeField := dstType.Field(i)
		valueField := dstValue.Elem().Field(i)
		jsonTag := ParseJsonTag(typeField.Tag.Get("json"))
		inline := (typeField.Anonymous && len(jsonTag.Name) == 0) || jsonTag.Inline
		name := jsonTag.Name
		if len(name) == 0 {
			name = typeField.Name
		}
		processingField = name
		var value interface{}
		var ok bool
		if !inline {
			value, ok = src[jsonTag.Name]
		} else {
			value, ok = src, true
		}
		if ok {
			if valueField.Kind() == reflect.Struct {
				if err = Merge(valueField.Addr().Interface(), value.(map[string]interface{})); err != nil {
					return
				}
			} else {
				convertedValue := reflect.ValueOf(value).Convert(typeField.Type)
				valueField.Set(convertedValue)
			}
		}
	}
	return
}
