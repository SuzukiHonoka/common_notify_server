package nfly

type GROUP int

const (
	// Admin are able to access any function
	Admin GROUP = iota
	// Common are able to access any common function
	Common GROUP = iota
	// Custom are able to access any custom allowed function
	Custom GROUP = iota
	// Baned can not access any function
	Baned GROUP = iota
)