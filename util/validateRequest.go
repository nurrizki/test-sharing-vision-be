package util

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ValidateRequest(param interface{}) error {
	types := reflect.TypeOf(param)
	values := reflect.ValueOf(param)
	if types.Kind() == reflect.Ptr || types.Kind() == reflect.Interface {
		types = types.Elem()
		values = values.Elem()
	}
	for i := 0; i < types.NumField(); i++ {
		value := values.Field(i)
		field := types.Field(i)
		validFuncs := field.Tag.Get("valid")
		if validFuncs == "" {
			continue
		}
		validName := field.Tag.Get("validname")
		if validName == "" {
			validName = field.Name
		}
		for _, validation := range strings.Split(validFuncs, ",") {
			validateDetail := strings.Split(validation, "=")
			if len(validateDetail) > 1 {
				// maximal
				if validateDetail[0] == "max" {
					max, _ := strconv.Atoi(validateDetail[1])
					if len(value.String()) > max {
						return fmt.Errorf("%s must not more than %d character", validName, max)
					}
				}
				// minimal
				if validateDetail[0] == "min" {
					min, _ := strconv.Atoi(validateDetail[1])
					if len(value.String()) < min {
						return fmt.Errorf("%s must not less than %d character", validName, min)
					}
				}
				// one of
				if validateDetail[0] == "oneof" {
					option := strings.Split(validateDetail[1], "|")

					mapOption := make(map[string]bool)
					for _, v := range option {
						mapOption[v] = true
					}
					if !mapOption[strings.TrimSpace(value.String())] {
						return fmt.Errorf("%s must be one of: %s", validName, strings.Join(option, ","))
					}
				}
			}

			if validation == "required" {
				switch value.Kind() {
				case reflect.String:
					if strings.TrimSpace(value.String()) == "" {
						return fmt.Errorf("%s must not be empty", validName)
					}
				case reflect.Int:
					if value.Int() == 0 {
						return fmt.Errorf("%s must not be 0", validName)
					}
				default:
					continue
				}
			}
		}
	}
	return nil
}
