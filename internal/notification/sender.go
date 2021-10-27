package notification

import "common_notify_server/internal/user"

type Sender struct {
	User *user.User
	// Worker
}
