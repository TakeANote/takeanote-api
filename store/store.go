package store

import (
	"time"

	"github.com/takeanote/takeanote-api/model"
	"golang.org/x/net/context"
)

// Store carries methods to manipulate data in a storage.
type Store interface {
	CreateUser(*model.User) error
	UpdateUser(*model.User) error
	GetUserByEmail(string) (*model.User, error)
	GetUserByEmailPassword(string, string) (*model.User, error)
	GetEmailByToken(string) (string, error)
	CreateToken(string, string, time.Duration) error
	DeleteToken(string) error
}

// CreateUser create a new user into the store.
func CreateUser(c context.Context, user *model.User) error {
	return FromContext(c).CreateUser(user)
}

// UpdateUser update an existing user into the store.
func UpdateUser(c context.Context, user *model.User) error {
	return FromContext(c).UpdateUser(user)
}

// GetUserByEmailPassword retrieve a user thanks to an email and password.
func GetUserByEmailPassword(c context.Context, email, password string) (*model.User, error) {
	return FromContext(c).GetUserByEmailPassword(email, password)
}

// GetUserByEmail retrieve a user thanks to an email.
func GetUserByEmail(c context.Context, email string) (*model.User, error) {
	return FromContext(c).GetUserByEmail(email)
}

// GetEmailByToken retrieve a user thanks to a token.
func GetEmailByToken(c context.Context, token string) (string, error) {
	return FromContext(c).GetEmailByToken(token)
}

// CreateToken insert a generated token into the store.
func CreateToken(c context.Context, token, email string, duration time.Duration) error {
	return FromContext(c).CreateToken(token, email, duration)
}

// DeleteToken delete the specified token from the store.
func DeleteToken(c context.Context, token string) error {
	return FromContext(c).DeleteToken(token)
}
