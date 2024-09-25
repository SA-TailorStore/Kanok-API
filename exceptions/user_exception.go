package exceptions

import "errors"

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrDuplicatedEmail = errors.New("duplicated email")
	ErrLoginFailed     = errors.New("login failed")
)
