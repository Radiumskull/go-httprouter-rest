package utils

import (
	"backend/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(user *models.User) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"] = user.Userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("WeLoveGito"))
	if err != nil {
		return "", err
	}
	return token, nil
}
