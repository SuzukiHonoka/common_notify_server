package nfly

// ID means the action IDs
type ID int

const (
	ACK ID = iota
	SYN ID = iota

	Interactive ID = iota
	Passive     ID = iota
)
