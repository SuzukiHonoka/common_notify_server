package notification

import (
	"github.com/google/uuid"
	"nfly/internal/user"
)

type Notification struct {
	Header  Header
	Message Message
	Status  Status
}

type Notifications []*Notification

type UserMap map[string]Notifications

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

func (x *UserMap) FindNotificationsByUser(u *user.User) Notifications {
	// filter status
	var notPushedNotifications Notifications
	for _, el := range (*x)[u.Credit.Email] {
		// check flag
		if !el.Status.Pushed {
			notPushedNotifications = append(notPushedNotifications, el)
		}
	}
	return notPushedNotifications
}

func (x *UserMap) DeleteNotificationsByUser(u *user.User) {
	delete(*x, u.Credit.Email)
}
