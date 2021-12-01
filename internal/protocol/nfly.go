package nfly

import "nfly/internal/notification"

type NFLY struct {
	Action   ID
	Type     notification.Priority
	Msg      notification.Message
	Receiver RECEIVER
}
