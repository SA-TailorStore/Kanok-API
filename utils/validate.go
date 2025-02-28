package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	valid "github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/liyue201/goqr"
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

func ValidateUsername(username string) error {
	re := regexp.MustCompile("^[a-zA-Z0-9_!@]+$")

	if !re.MatchString(username) {
		return exceptions.ErrInvalidFormatUsername
	}
	return nil
}

func ValidatePassword(pass string) error {
	if len(pass) < 8 {
		return exceptions.ErrCharLeastPassword
	}

	re := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)
	if !re.MatchString(pass) {
		return exceptions.ErrOneSpecialPassword
	}

	return nil
}

func ValidatePhoneNumber(phone string) error {
	if len(phone) < 10 {
		return exceptions.ErrLeastPhoneNumber
	}

	re := regexp.MustCompile(`^0[2689][0-9]{8}$`)
	if !re.MatchString(phone) {
		return exceptions.ErrPhoneNumber
	}
	return nil
}

func ValidateJWTFormat(tokenString string) error {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return exceptions.ErrInvalidToken
	}

	_, _, err := jwt.NewParser().ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return fmt.Errorf("invalid JWT: %v", err)
	}

	return nil
}

func ValidateSlip(qrcode []*goqr.QRData) error {

	for _, code := range qrcode {
		parsedCode := ParseCode(string(code.Payload))
		data := strings.ReplaceAll(string(code.Payload), parsedCode.Checksum, "")
		checksumqr, _ := CRC16XModem(data, 0xffff)
		checksum, err := HexToDecimal(parsedCode.Checksum)

		if err != nil {
			return err
		}

		if checksumqr != checksum {
			return exceptions.ErrWrongSlip
		}
	}

	return nil
}

func ValidateName(o string, s string) error {

	if o == s {
		return exceptions.ErrDupicatedName
	}

	return nil
}
