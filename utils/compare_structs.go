package utils

import "reflect"

func CompareStructs(a, b interface{}) bool {
	aVal := reflect.ValueOf(a)
	bVal := reflect.ValueOf(b)

	if aVal.Kind() != reflect.Struct || bVal.Kind() != reflect.Struct {
		return false
	}

	for i := 0; i < aVal.NumField(); i++ {
		aField := aVal.Field(i)
		bField := bVal.Field(i)

		if aField.Type() != bField.Type() {
			return false
		}

		switch aField.Kind() {
		case reflect.Struct:
			if !CompareStructs(aField.Interface(), bField.Interface()) {
				return false
			}
		case reflect.Ptr:
			if !aField.IsNil() && !bField.IsNil() {
				if !CompareStructs(aField.Elem().Interface(), bField.Elem().Interface()) {
					return false
				}
			} else if aField.IsNil() != bField.IsNil() {
				return false
			}
		default:
			if !reflect.DeepEqual(aField.Interface(), bField.Interface()) {
				return false
			}
		}
	}
	return true
}
