package gomerge

import (
	"fmt"
	"reflect"
)

func Merge(dst interface{}, src map[string]interface{}) (err error) {
	var processing_field string
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s has wrong value type: %s", processing_field, r)
		}
	}()
	dst_value := reflect.ValueOf(dst)
	dst_type := dst_value.Elem().Type()
	for i := 0; i < dst_type.NumField(); i++ {
		type_field := dst_type.Field(i)
		value_field := dst_value.Elem().Field(i)
		json_tag := type_field.Tag.Get("json")
		processing_field = json_tag
		if value, ok := src[json_tag]; ok {
			converted_value := reflect.ValueOf(value).Convert(type_field.Type)
			value_field.Set(converted_value)
		}
	}
	return
}
