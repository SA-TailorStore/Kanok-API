package exceptions

import "errors"

var (
	ErrMaterialNotFound   = errors.New("material not found")
	ErrBadRequestMaterial = errors.New("bad request data type wrong")
)
