package main

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	Name string `validate:"required"`
	Age  int    `validate:"min=18,max=60"`
}

func validateStruct(obj interface{}) error {
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		fieldName := strings.Trim(fieldType.Name, " ")

		fieldValue := field.Interface()
		if field.Type().Kind() == reflect.String {
			fieldValue = strings.TrimSpace(field.String())
		}

		// Check for required field
		if fieldType.Tag.Get("validate") == "required" {
			if fieldValue == reflect.Zero(field.Type()).Interface() {
				return fmt.Errorf("Field %s is required", fieldName)
			}
		}
	}

	return nil
}

func main() {
	user := User{
		Name: "     ",
		Age:  200,
	}

	// Name := strings.TrimSpace("   ")

	// if Name != "" {
	// 	fmt.Println(Name)
	// }

	err := validateStruct(user)
	if err != nil {
		fmt.Println("Validation failed:", err)
	} else {
		fmt.Println("Berhasil Dibuat")
	}
}
