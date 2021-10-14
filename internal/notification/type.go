package notification

type Type uint8

const (
	TypeSenderAPI    Type = iota
	TypeSenderWorker Type = iota
)
