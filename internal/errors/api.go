package errors

import "errors"

var (
	SessionPoolMaxReached = errors.New("session pool max reached")
	ParseAccountFailed    = errors.New("parse account failed")
)
