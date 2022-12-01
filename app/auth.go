package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/irfan44/task-5-vix-btpns-IrfanNurghiffariM/models"
)

var jwtKey = []byte("my_secret_key")

func GenerateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"id":       user.ID,
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func GetEmailFromToken(tokenString string) (email string, err error) {
	token, err := ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return "", err
	}

	email = (*claims)["email"].(string)

	return email, nil
}
