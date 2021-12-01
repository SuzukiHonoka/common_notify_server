package nfly

import (
	"net"
	client "nfly/internal/client"
)

type RECEIVER struct {
	Conn   net.Conn
	Client client.Client
}
