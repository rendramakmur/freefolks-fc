package helper

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	UserId   int    `json:"userId"`
	Email    string `json:"email"`
	UserType int    `json:"userType"`
	jwt.RegisteredClaims
}

func GenerateToken(jc *JwtClaims) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	newClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jc)
	ss, err := newClaims.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return ss, nil
}

func ParseToken(tokenString string) (*JwtClaims, error) {
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Invalid Token")
}
