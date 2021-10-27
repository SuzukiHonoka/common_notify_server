package notification

import (
	"common_notify_server/internal/user"
	"github.com/google/uuid"
)

type Notification struct {
	Header  Header
	Message Message
	Status  Status
}

type Notifications []*Notification

type UserMap map[*user.User]Notifications

func NewNotification(u *user.User, title *string, chain MessageChain) *Notification {
	uid, _ := uuid.NewRandom()
	return &Notification{
		Header: Header{
			UUID:     uid,
			Priority: PriorityMax, // set to max as default now
			Sender: Sender{
				User: u,
			},
		},
		Message: Message{
			Title:        title,
			MessageChain: chain,
		},
		Status: Status{
			Pushed: false,
		},
	}
}

func (x UserMap) FindNotificationsByUser(u *user.User) []*Notification {
	// filter status
	var notPushedNotifications Notifications
	for _, el := range x[u] {
		// check flag
		if !el.Status.Pushed {
			notPushedNotifications = append(notPushedNotifications, el)
		}
	}
	return notPushedNotifications
}
