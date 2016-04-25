package datastore

import "github.com/takeanote/takeanote-api/model"

// CreateUser creates a new user.
func (ds *datastore) CreateUser(user *model.User) error {
	return ds.Save(user).Error
}

// CreateUser creates a new user.
func (ds *datastore) UpdateUser(user *model.User) error {
	return ds.Save(user).Error
}

// GetUserByEmailPassword retrieve a user thanks to an email and password.
func (ds *datastore) GetUserByEmailPassword(email, password string) (*model.User, error) {
	var user model.User
	err := ds.Where("email = ? AND password = ?", email, password).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail retrieve a user thanks to an email.
func (ds *datastore) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := ds.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
