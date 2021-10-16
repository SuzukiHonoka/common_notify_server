package notification

import "common_notify_server/internal/user"

type Notification struct {
	Title    *string
	Message  []*interface{}
	Priority Priority
	Type     Type
	Data     []byte
	Status
	// addition files
}

func FindNotificationsByUser(u *user.USER) []*Notification {
	return CachedNotifications[u]
}
