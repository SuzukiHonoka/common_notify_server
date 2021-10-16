package notification

import "common_notify_server/internal/user"

var CachedNotifications map[*user.USER][]*Notification
