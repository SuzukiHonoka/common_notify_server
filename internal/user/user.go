package user

import (
	userErrors "common_notify_server/internal/errors"
	iface "common_notify_server/internal/interface"
	"golang.org/x/crypto/bcrypt"
)

var Helper iface.Helper

type User struct {
	Credit Credit
	Group  Group
	// Limit
	// Access Control
}

func (u *User) ComparePassword(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Credit.Password), []byte(pass)) == nil
}

// Register the user to the database and return User instance
// default group is Admin
func Register(email string, pass string, group *Group) (*User, error) {
	// if user exist
	if userExist(email) {
		return nil, userErrors.UserExist
	}
	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	// return error if hash failed
	if err != nil {
		return nil, err
	}
	// using default if nil group
	if group == nil {
		group = AdminGP
	}
	// create an instance of User
	n := &User{
		Credit: Credit{
			Email:    email,
			Password: string(hash),
		},
		Group: *group,
	}
	// store to cache and DB
	if addNewUser(n) {
		return n, nil
	}
	// return error if store process failed
	return nil, userErrors.UserSaveFailed
}

// Login the user and return User instance
func Login(email string, pass string) (*User, error) {
	// find the user
	u := findUserByEmail(email)
	// return error if not found by id or email
	if u == nil {
		return nil, userErrors.UserNotFound
	}
	// password authentication
	if u.ComparePassword(pass) {
		return u, nil
	}
	// return error if authentication failed
	return nil, userErrors.UserAuthenticationFailed
}

// Refresh the cached User slice from DB
func Refresh() bool {
	Helper.Refresh()
	return true
}

// addNewUser to cached User slice and DB
func addNewUser(user *User) bool {
	CachedUsers = append(CachedUsers, user)
	// save users to DB
	Helper.AddUser(user)
	return true
}

func userExist(email string) bool {
	return findUserByEmail(email) != nil
}

func findUserByEmail(email string) *User {
	for _, user := range CachedUsers {
		if user.Credit.Email == email {
			return user
		}
	}
	return nil
}
