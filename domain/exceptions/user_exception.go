package exceptions

import "errors"

var (
	ErrUserNotFound           = errors.New("user not found")
	ErrUsernameDuplicated     = errors.New("duplicated username")
	ErrUserPasswordFormat     = errors.New("password format is invalid")
	ErrUsernameSameAsPassword = errors.New("username and password must be different")
	ErrUsernameFormat         = errors.New("username format is invalid")
	ErrInvalidPassword        = errors.New("password must be at least 8 characters long")
	ErrLoginFailed            = errors.New("login failed")
	ErrInvalidToken           = errors.New("invalid token")
	ErrExpiredToken           = errors.New("token is expired")
)
