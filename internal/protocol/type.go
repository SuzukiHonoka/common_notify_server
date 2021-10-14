package nfly

// TYPE means the notice type by priority
type TYPE int

const (
	COMMON TYPE = iota
	Vital  TYPE = iota
	// opt: bar only, fullscreen..
)
