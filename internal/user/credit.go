package user

type Credit struct {
	Email    string
	Password string `json:"-"` // hide hashed password string from marshall
}
