package errors

import "errors"

var (
	SessionPoolMaxReached = errors.New("session pool max reached")
)
