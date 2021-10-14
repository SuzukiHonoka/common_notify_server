package notification

type Priority uint8

const (
	PriorityMax    Priority = iota           // popup until user dismiss it
	PriorityHigh   Priority = iota           // popup with long time
	PriorityMedium Priority = iota           // popup with short time
	PriorityLow    Priority = iota           // never popup
	PriorityCommon Priority = PriorityMedium // eq to medium
)
