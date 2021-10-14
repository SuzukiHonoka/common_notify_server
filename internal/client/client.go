package client

import (
	"common_notify_server/internal/user"
	"github.com/google/uuid"
)

type Client struct {
	ID         uuid.UUID
	DeviceInfo Device
	Bond       []user.USER
	Version    string
}
