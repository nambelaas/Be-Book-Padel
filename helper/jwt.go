package helper

import (
	"Be-Book-Padel/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var jwtSecret = []byte(viper.GetString("JWT_SECRET"))

type JwtData struct {
	jwt.RegisteredClaims
	UserId uint
	Name   string
	Role   string
}

func GenerateAccessToken(userID uint, role models.UserRole) (string, error) {
	claims := jwt.MapClaims{
		"user_id": uint(userID),
		"role":    role,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (JwtData, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtData{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return JwtData{}, err
	}

	claims, ok := token.Claims.(*JwtData)
	if !ok {
		return JwtData{}, err
	}

	return *claims, nil
}
