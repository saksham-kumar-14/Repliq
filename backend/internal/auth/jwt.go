package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/saksham-kumar-14/Repliq/backend/internal/env"
)

var jwtSecret = []byte(env.GetString("JWT_SECRET", "supersecret"))

func GenerateJWT(userID uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseJWT(token_string string) (*jwt.Token, error) {
	return jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
}
