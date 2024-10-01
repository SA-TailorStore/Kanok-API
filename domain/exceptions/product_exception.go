package exceptions

import "errors"

var (
	ErrProductNotFound = errors.New("product not found")
	ErrDupicatedID     = errors.New("ID dupicated")
)
