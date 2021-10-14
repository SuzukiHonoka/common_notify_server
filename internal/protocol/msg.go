package nfly

import "time"

type MSG struct {
	Title   string
	Content string
	// Source Worker
	Remark    string
	Timestamp time.Time
}
