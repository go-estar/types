package structUtil

import (
	"github.com/pkg/errors"
	"reflect"
)

var ErrorFieldInvalid = "field %s invalid"
var ErrorFieldCantSet = "field %s can't set"

func SetValue(obj interface{}, field string, value interface{}) error {
	valueV := reflect.ValueOf(obj)
	if valueV.Kind() == reflect.Ptr {
		valueV = valueV.Elem()
	}

	fieldV := valueV.FieldByName(field)
	if !fieldV.IsValid() {
		return errors.Errorf(ErrorFieldInvalid, field)
	}

	if !fieldV.CanSet() {
		return errors.Errorf(ErrorFieldCantSet, field)
	}

	fieldV.Set(reflect.ValueOf(value))
	return nil
}

func GetValue(obj interface{}, field string) (interface{}, error) {
	valueV := reflect.ValueOf(obj)
	if valueV.Kind() == reflect.Ptr {
		valueV = valueV.Elem()
	}

	fieldV := valueV.FieldByName(field)
	if !fieldV.IsValid() {
		return nil, errors.Errorf(ErrorFieldInvalid, field)
	}
	return fieldV.Interface(), nil
}
