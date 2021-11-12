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

type UsersMap map[string]*User
type UsersList []*User

func (u *User) ComparePassword(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Credit.Password), []byte(pass)) == nil
}

// Register the user to the database and return User instance
// default group is Admin
func Register(email string, pass string, group *Group) (*User, error) {
	// if user exist
	if CachedUsersMap.UserExist(email) {
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
	if CachedUsersMap.AddNewUser(n) {
		return n, nil
	}
	// return error if store process failed
	return nil, userErrors.UserSaveFailed
}

// Login the user and return User instance
func Login(email string, pass string) (*User, error) {
	// find the user
	u := CachedUsersMap.FindUserByEmail(email)
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

// AddNewUser to cached User slice and DB
func (x *UsersMap) AddNewUser(user *User) bool {
	(*x)[user.Credit.Email] = user
	// save users to DB
	Helper.AddUser(user)
	return true
}

// DeleteUser cached User from slice and DB
func (x *UsersMap) DeleteUser(user *User) bool {
	// find and clean
	delete(*x, user.Credit.Email)
	Helper.DelUser(user)
	return true
}

func (x *UsersMap) UserExist(email string) bool {
	return x.FindUserByEmail(email) != nil
}

func (x *UsersMap) FindUserByEmail(email string) *User {
	if v, ok := (*x)[email]; ok {
		return v
	}
	return nil
}
