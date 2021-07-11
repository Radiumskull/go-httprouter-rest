package utils

import (
	"backend/models"
	"crypto/ed25519"

	"github.com/pascaldekloe/jwt"
)

func ParseJWT(token string) (*jwt.Claims, error) {
	var JWTPrivateKey = "WeLoveGito"
	claims, err := jwt.EdDSACheck([]byte(token), ed25519.PublicKey([]byte(JWTPrivateKey)))
	return claims, err
}

func EncodeJWT(user *models.User) (string, error) {
	var (
		claims        jwt.Claims
		JWTPrivateKey = "WeLoveGito"
	)

	claims.ID = string(rune(user.Userid))
	claims.Subject = user.Username

	token, err := claims.EdDSASign(ed25519.PrivateKey([]byte(JWTPrivateKey)))

	return string(token), err
}
