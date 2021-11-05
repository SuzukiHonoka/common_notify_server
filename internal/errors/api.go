package errors

import "errors"

var (
	BadRequest                   = errors.New("bad request")
	SessionPoolMaxReached        = errors.New("session pool max reached")
	SessionInvalid               = errors.New("session invalid")
	NotificationsListParseFailed = errors.New("notifications list parse failed")
	ParseAccountFailed           = errors.New("parse account failed")
)
