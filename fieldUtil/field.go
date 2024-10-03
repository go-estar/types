package fieldUtil

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func IsNil(s interface{}) bool {
	if s == nil {
		return true
	}
	v := reflect.ValueOf(s)
	return v.Kind() == reflect.Ptr && v.IsNil()
}

func IsEmpty(s interface{}) bool {
	if s == nil {
		return true
	}
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return true
		}
		return v.Elem().IsZero()
	} else if v.Kind() == reflect.Map || v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		return v.Len() == 0
	}
	return v.IsZero()
}

func GetValue(field reflect.Value) string {
	if field.Kind() == reflect.Map || field.Kind() == reflect.Slice || field.Kind() == reflect.Array || field.Kind() == reflect.Struct {
		if (field.Kind() == reflect.Map || field.Kind() == reflect.Slice || field.Kind() == reflect.Array) && field.Len() == 0 {
			return ""
		}
		if field.Kind() == reflect.Struct && field.IsZero() {
			return ""
		}
		byteVal, err := json.Marshal(field.Interface())
		if err == nil {
			return string(byteVal)
		} else {
			return fmt.Sprint(field.Interface())
		}
	} else {
		return fmt.Sprint(field.Interface())
	}
}

func GetValueArr(field reflect.Value) []string {
	var value []string
	if field.Kind() == reflect.Map || field.Kind() == reflect.Slice || field.Kind() == reflect.Array || field.Kind() == reflect.Struct {
		if (field.Kind() == reflect.Map || field.Kind() == reflect.Slice || field.Kind() == reflect.Array) && field.Len() == 0 {
			return value
		}
		if field.Kind() == reflect.Struct && field.IsZero() {
			return value
		}
		if field.Kind() == reflect.Slice || field.Kind() == reflect.Array {
			for i := 0; i < field.Len(); i++ {
				value = append(value, fmt.Sprint(field.Index(i).Interface()))
			}
		} else {
			byteVal, err := json.Marshal(field.Interface())
			if err == nil {
				value = append(value, string(byteVal))
			} else {
				value = append(value, fmt.Sprint(field.Interface()))
			}
		}
	} else {
		value = append(value, fmt.Sprint(field.Interface()))
	}
	return value
}
