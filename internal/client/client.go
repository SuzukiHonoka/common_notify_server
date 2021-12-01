package client

import (
	"github.com/google/uuid"
	"nfly/internal/user"
)

type Client struct {
	ID         uuid.UUID
	DeviceInfo Device
	Bond       []*user.User
	Version    string
}
