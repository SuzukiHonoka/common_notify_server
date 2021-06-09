package nfly

type PERMISSION int

const (
	// User represents the User Management
	User PERMISSION = iota
	// API represents the API Management
	API PERMISSION = iota
	// Worker represents the Worker Management
	Worker PERMISSION = iota
	// SafeGuard represents the SafeGuard Management
	SafeGuard PERMISSION = iota
	// Generator represents the Generator Management
	Generator PERMISSION = iota
)
