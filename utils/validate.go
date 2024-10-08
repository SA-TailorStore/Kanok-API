package utils

import (
	"fmt"
	"strings"

	valid "github.com/go-playground/validator/v10"
)

var validate = valid.New()

type ValidateError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func ValidateStruct[T any](payload T) *ValidateError {
	err := validate.Struct(payload)

	if err != nil {
		// Handle ValidationErrors
		errMsg := ""
		if validationErrors, ok := err.(valid.ValidationErrors); ok {
			for _, err := range validationErrors {
				tmp := strings.Split(err.StructNamespace(), ".")
				msg := fmt.Sprintf("%s is %s", tmp[len(tmp)-1], err.Tag())
				msg = strings.ToLower(string(msg[0])) + msg[1:] // lowercase the first letter
				errMsg = errMsg + msg + ", "
			}

			return &ValidateError{
				Error:   "Invalid request",
				Message: errMsg[:len(errMsg)-2], // Remove the trailing comma and space
			}
		}

		// Return a generic error message if we encounter an unknown error
		return &ValidateError{
			Error:   "Unknown validation error",
			Message: err.Error(),
		}
	}

	return nil
}

func ValidateUsername(u string) *ValidateError {

	return nil
}
func ValidatePassword(u string) *ValidateError {

	return nil
}
