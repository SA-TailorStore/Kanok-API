package exceptions

import "errors"

var (
	ErrUserNotFound           = errors.New("user not found")
	ErrUsernameDuplicated     = errors.New("duplicated username")
	ErrUserPasswordFormat     = errors.New("password format is invalid")
	ErrUsernameSameAsPassword = errors.New("username and password must be different")
	ErrInvalidFormatUsername  = errors.New("username format is invalid")
	ErrWrongUsername          = errors.New("username is wrong ")
	ErrWrongPassword          = errors.New("password is wrong ")
	ErrCharLeastPassword      = errors.New("password must be at least 8 characters long")
	ErrOneSpecialPassword     = errors.New("password must contain at least one special character")
	ErrLoginFailed            = errors.New("login failed")
	ErrInvalidToken           = errors.New("invalid token")
	ErrExpiredToken           = errors.New("token is expired")
	ErrPhoneNumber            = errors.New("invalid phone number")
	ErrLeastPhoneNumber       = errors.New("phone number must be at least 10")
	ErrRoleNotHave            = errors.New("not have this role")
	ErrInvalidImage           = errors.New("slip is wrong")
	ErrNoImage                = errors.New("no image")
)
