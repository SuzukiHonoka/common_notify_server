package nfly

import "common_notify_server/internal/notification"

type NFLY struct {
	Action   ID
	Type     notification.Priority
	Msg      notification.Message
	Receiver RECEIVER
}
