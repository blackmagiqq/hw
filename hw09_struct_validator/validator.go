package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var errMsgs []string
	for _, ve := range v {
		errMsgs = append(errMsgs, fmt.Sprintf("%s: %v", ve.Field, ve.Err))
	}
	return strings.Join(errMsgs, ", ")
}

func Validate(v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Struct {
		return errors.New("input is not a struct")
	}

	var validationErrors ValidationErrors

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		fieldValue := val.Field(i)
		tag := field.Tag.Get("validate")

		if tag == "" {
			continue
		}

		validators := strings.Split(tag, "|")
		if fieldValue.Kind() == reflect.Slice {
			for j := 0; j < fieldValue.Len(); j++ {
				elemValue := fieldValue.Index(j)
				for _, validator := range validators {
					if err := applyValidator(field.Name, elemValue, validator); err != nil {
						validationErrors = append(validationErrors, ValidationError{Field: field.Name, Err: err})
					}
				}
			}
		} else {
			for _, validator := range validators {
				if err := applyValidator(field.Name, fieldValue, validator); err != nil {
					validationErrors = append(validationErrors, ValidationError{Field: field.Name, Err: err})
				}
			}
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}
	return nil
}

func applyValidator(fieldName string, fieldValue reflect.Value, validator string) error {
	parts := strings.SplitN(validator, ":", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid validator format: %s", validator)
	}

	switch parts[0] {
	case "len":
		if fieldValue.Kind() == reflect.String {
			expectedLen, _ := strconv.Atoi(parts[1])
			if len(fieldValue.String()) != expectedLen {
				return fmt.Errorf("length must be %d", expectedLen)
			}
		}
	case "regexp":
		if fieldValue.Kind() == reflect.String {
			re, err := regexp.Compile(parts[1])
			if err != nil {
				return fmt.Errorf("invalid regexp: %s", parts[1])
			}
			if !re.MatchString(fieldValue.String()) {
				return fmt.Errorf("must match regexp %s", parts[1])
			}
		}
	case "in":
		options := strings.Split(parts[1], ",")
		switch fieldValue.Kind() {
		case reflect.String:
			if !contains(options, fieldValue.String()) {
				return fmt.Errorf("must be one of %v", options)
			}
		case reflect.Int:
			intOptions := make([]int, len(options))
			for i, opt := range options {
				intOptions[i], _ = strconv.Atoi(opt)
			}
			if !containsInt(intOptions, int(fieldValue.Int())) {
				return fmt.Errorf("must be one of %v", intOptions)
			}
		}
	case "min":
		if fieldValue.Kind() == reflect.Int {
			minValue, _ := strconv.Atoi(parts[1])
			if int(fieldValue.Int()) < minValue {
				return fmt.Errorf("must be at least %d", minValue)
			}
		}
	case "max":
		if fieldValue.Kind() == reflect.Int {
			maxValue, _ := strconv.Atoi(parts[1])
			if int(fieldValue.Int()) > maxValue {
				return fmt.Errorf("must be at most %d", maxValue)
			}
		}
	}

	return nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func containsInt(slice []int, item int) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
