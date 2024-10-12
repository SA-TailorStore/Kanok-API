package exceptions

import "errors"

var (
	ErrUploadImage = errors.New("image upload failed")
)
