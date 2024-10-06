package exceptions

import "errors"

var (
	ErrProductNotFound    = errors.New("product not found")
	ErrDupicatedProductID = errors.New("ID dupicated")
)
