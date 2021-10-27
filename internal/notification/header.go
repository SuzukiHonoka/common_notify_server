package notification

import (
	"github.com/google/uuid"
)

type Header struct {
	UUID     uuid.UUID
	Priority Priority
	Sender   Sender
}
