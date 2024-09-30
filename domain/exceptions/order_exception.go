package exceptions

import "errors"

var (
	ErrOrderNotFound = errors.New("order not found")
	ErrInfomation    = errors.New("infomation not found")
)
