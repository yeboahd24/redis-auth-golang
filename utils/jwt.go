package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// Claims defines the JWT claims.
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// GenerateJWT creates a JWT token for a given email.
func GenerateJWT(email, secretKey string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
