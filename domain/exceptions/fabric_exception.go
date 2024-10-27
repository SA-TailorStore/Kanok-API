package exceptions

import "errors"

var (
	ErrFabricNotFound  = errors.New("fabric not found")
	ErrFabricNotEnough = errors.New("fabric not enough")
)
