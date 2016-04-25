package datastore

import "time"

// GetEmailByToken retrieve a user thanks to a token.
func (ds *datastore) GetEmailByToken(token string) (string, error) {
	return ds.RedisClient.Get(token).Result()
}

// CreateToken insert a generated token into the store.
func (ds *datastore) CreateToken(token, email string, duration time.Duration) error {
	return ds.RedisClient.Set(token, email, duration).Err()
}

// DeleteToken delete the specified token from the store.
func (ds *datastore) DeleteToken(token string) error {
	return ds.RedisClient.Del(token).Err()
}
