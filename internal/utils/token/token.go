package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("cwMDQsImlhdCI6MTc0NTE4NzgwNCwic3ViIjoiMzhlNTczN2QtM")

const tokenExpiry = time.Hour * 72

func GenerateJWT(accountID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": accountID,
		"exp": time.Now().Add(tokenExpiry).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
