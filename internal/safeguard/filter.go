package safeguard

import "net"

type Action int

const (
	Block  Action = iota
	Accept Action = iota
)

type Filter struct {
	ID     uint8
	Name   string
	Action Action
	IPs    []net.IP
}
