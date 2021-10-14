package user

type Type int

type Group struct {
	ID         uint8
	Name       string
	Type       Type
	Permission *[]Permission // only require if Type is Custom
}

const (
	// Admin are able to access any function
	Admin Type = iota
	// Common are able to access any common function
	Common Type = iota
	// Custom are able to access any custom allowed function
	Custom Type = iota
	// Baned can not access any function
	Baned Type = iota
)

var (
	AdminGP = &Group{
		ID:   0,
		Name: "管理员",
		Type: Admin,
	}
)
