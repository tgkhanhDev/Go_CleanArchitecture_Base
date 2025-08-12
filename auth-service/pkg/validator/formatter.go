package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(err error) map[string]string {
	// Ép kiểu lỗi sang validator.ValidationErrors
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		// handle another if neccessary
		return nil
	}

	errors := make(map[string]string)

	for _, fieldErr := range validationErrors {
		field := fieldErr.Field()
		tag := fieldErr.Tag()
		param := fieldErr.Param()

		switch tag {
		case "required":
			errors[field] = fmt.Sprintf("The %s field is required.", field)
		case "email":
			errors[field] = fmt.Sprintf("The %s field must be a valid email address.", field)
		case "min":
			errors[field] = fmt.Sprintf("The %s field must be at least %s characters long.", field, param)
		case "max":
			errors[field] = fmt.Sprintf("The %s field must not exceed %s characters.", field, param)
		default:
			errors[field] = fmt.Sprintf("The %s field is invalid.", field)
		}
	}

	return errors
}

func ErrorBuilder(err map[string]string) string {
	if len(err) == 0 {
		return ""
	}

	errMsg := "Validation errors:\n"
	for field, message := range err {
		errMsg += fmt.Sprintf("%s: %s\n", field, message)
	}
	return errMsg
}
