package user

type Permission int

const (
	// User represents the User Management
	User Permission = iota
	// API represents the API Management
	API Permission = iota
	// Worker represents the Worker Management
	Worker Permission = iota
	// SafeGuard represents the SafeGuard Management
	SafeGuard Permission = iota
	// Generator represents the Generator Management
	Generator Permission = iota
)
