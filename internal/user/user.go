package nfly

type USER struct {
	Name string
	Group GROUP
	Permission []PERMISSION
	// Limit
	// Access Control
}

// Register the user into the database and return USER instance
func Register(name string, pass string) *USER{
	return &USER{}
}

