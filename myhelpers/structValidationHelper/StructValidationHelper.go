package structValidationHelper

import (
	"errors"
	"fmt"
	"reflect"

	set "hcn/myhelpers/setHelper"
)

// ValidateFields is an aux function that checks if the fields of the struct
// are not empty and has the correct data type.
// structFields is an array of fields to check
func ValidateFields(model interface{}, structFields ...[]string) (bool, error) {
	// sometimes we only have to check some fields
	fields := set.New()
	if len(structFields) != 0 {
		for _, field := range structFields[0] {
			fields.Insert(field)
		}
	}

	v := reflect.ValueOf(model)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := v.Type().Field(i).Name

		// only checks the fields requested
		if fields.Has(fieldName) || len(structFields) == 0 {
			fieldType := fmt.Sprintf("%T", field.Interface())
			fieldValue := field.Elem()
			//fmt.Println("Name:", fieldName, "     Type:", fieldType, "     FieldValue:", fieldValue)
			if fieldValue.IsValid() {
				switch fieldType {
				case "*int":
					value := fmt.Sprintf("%v", field.Elem()) // we have to extract the value this way for the if statement
					numberSign := fmt.Sprintf("%c", value[0])
					if value == "0" || numberSign == "-" {
						return false, errors.New(fieldName + " is empty or not valid")
					}
				case "*string":
					value := fmt.Sprintf("%v", field.Elem())
					if len(value) == 0 {
						return false, errors.New(fieldName + " is empty or not valid")
					}
				}
			} else {
				return false, errors.New(fieldName + " is empty or not valid")
			}
		}
	}
	return true, nil
}

// GetFields is an aux function that returns struct name, fields names and
// field values of a model
func GetFields(model interface{}) (string, []string, []string, error) {
	var structName string
	var fieldNames []string
	var fieldValues []string

	v := reflect.ValueOf(model)
	structName = reflect.TypeOf(model).Name()
	for i := 0; i < v.NumField(); i++ {

		field := v.Field(i)
		fieldName := v.Type().Field(i).Name
		fieldValue := fmt.Sprintf("%v", field.Elem())

		fieldNames = append(fieldNames, fieldName)
		fieldValues = append(fieldValues, fieldValue)
	}
	return structName, fieldNames, fieldValues, nil
}
