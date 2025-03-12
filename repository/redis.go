package repository

import (
	"fmt"

	"redis-auth/config"
	"redis-auth/models"

	"golang.org/x/crypto/bcrypt"
)

// GetUserByEmail retrieves the hashed password for a given email.
func GetUserByEmail(email string) (string, error) {
	key := fmt.Sprintf("user:%s", email)
	return config.RedisClient.Get(config.Ctx, key).Result()
}

// CreateUser stores a new user with a hashed password.
func CreateUser(user models.User) error {
	key := fmt.Sprintf("user:%s", user.Email)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return config.RedisClient.Set(config.Ctx, key, hashedPassword, 0).Err()
}

// UserExists checks if a user already exists.
func UserExists(email string) (bool, error) {
	key := fmt.Sprintf("user:%s", email)
	exists, err := config.RedisClient.Exists(config.Ctx, key).Result()
	return exists == 1, err
}
