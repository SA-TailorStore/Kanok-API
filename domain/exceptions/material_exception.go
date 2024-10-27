package exceptions

import "errors"

var (
	ErrMaterialNotFound   = errors.New("material not found")
	ErrBadRequestMaterial = errors.New("bad request data type wrong")
	ErrDupicatedName      = errors.New("this name is dupicated")
)
