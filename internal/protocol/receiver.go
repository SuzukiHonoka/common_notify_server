package nfly

import (
	client "common_notify_server/internal/client"
	"net"
)

type RECEIVER struct {
	Conn   net.Conn
	Client client.Client
}
