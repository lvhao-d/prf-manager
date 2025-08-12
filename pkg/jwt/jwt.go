package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("secret")

func GenerateToken(userID int, username string) (string, error) {
	claim := jwt.MapClaims{

		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
