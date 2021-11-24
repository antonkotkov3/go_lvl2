package main

import (
	"fmt"
	"reflect"
)

func FillStruct(in interface{}, data map[string]interface{}) error {
	for key, value := range data {

		err := setField(in, key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func setField(obj interface{}, name string, value interface{}) error {
	if obj == nil {
		return fmt.Errorf("Empty input object")
	}

	structValue := reflect.ValueOf(obj)
	if structValue.Kind() == reflect.Ptr {
		structValue = structValue.Elem()
	}
	fieldVal := structValue.FieldByName(name)

	if !fieldVal.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !fieldVal.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	val := reflect.ValueOf(value)

	if fieldVal.Type() == val.Type() {
		fieldVal.Set(val)
		return nil
	}

	if m, ok := value.(map[string]interface{}); ok {
		if fieldVal.Kind() == reflect.Struct {
			return FillStruct(fieldVal.Addr().Interface(), m)
		}
		if fieldVal.Kind() == reflect.Ptr && fieldVal.Type().Elem().Kind() == reflect.Struct {
			if fieldVal.IsNil() {
				fieldVal.Set(reflect.New(fieldVal.Type().Elem()))
			}
			return FillStruct(fieldVal.Interface(), m)
		}

	}
	return fmt.Errorf("Provided value type didn't match obj field type")
}
