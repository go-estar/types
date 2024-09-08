package structUtil

import (
	stderrors "errors"
	"github.com/pkg/errors"
	"reflect"
)

var ErrorFieldInvalid = stderrors.New("field invalid")
var ErrorFieldCantSet = stderrors.New("field can't set")

func SetValue(obj interface{}, field string, value interface{}) error {
	valueV := reflect.ValueOf(obj)
	if valueV.Kind() == reflect.Ptr {
		valueV = valueV.Elem()
	}

	fieldV := valueV.FieldByName(field)
	if !fieldV.IsValid() {
		return errors.WithMessage(ErrorFieldInvalid, field)
	}

	if !fieldV.CanSet() {
		return errors.WithMessage(ErrorFieldCantSet, field)
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
		return nil, errors.WithMessage(ErrorFieldInvalid, field)
	}
	return fieldV.Interface(), nil
}
