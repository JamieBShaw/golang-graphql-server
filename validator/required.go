package validator

import (
	"fmt"
	"reflect"
)

func (v *Validator) Required(field string, value interface{}) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}
	if isEmpty(value) {
		v.Errors[field] = fmt.Sprintf("%s is required", field)
	}
	return true
}

func isEmpty(value interface{}) bool {
	t := reflect.ValueOf(value)

	switch t.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		return t.Len() == 0
	}
	return false
}
