package exceptions

import "errors"

var (
	ErrUserNotFound           = errors.New("user not found")
	ErrDuplicatedUsername     = errors.New("duplicated username")
	ErrPasswordFormat         = errors.New("password format is invalid")
	ErrUsernameSameAsPassword = errors.New("username and password must be different")
	ErrUsernameFormat         = errors.New("username format is invalid")
	ErrInvalidPassword        = errors.New("password must be at least 8 characters long")
	ErrLoginFailed            = errors.New("login failed")
)
