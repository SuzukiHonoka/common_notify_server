package errors

import "errors"

var (
	SessionPoolMaxReached        = errors.New("session pool max reached")
	SessionInvalid               = errors.New("session invalid")
	NotificationsListParseFailed = errors.New("notifications list parse failed")
	ParseAccountFailed           = errors.New("parse account failed")
)
