package errors

import "errors"

var (
	UserExist                = errors.New("user exist")
	UserNotFound             = errors.New("user not found")
	UserAuthenticationFailed = errors.New("user authentication failed")
	UserSaveFailed           = errors.New("user save failed")
)
