package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type JwtData struct {
	jwt.RegisteredClaims
	UserId uint
	Name   string
	Role   string
}

func GetJwtTokenFromHeader(c *gin.Context) (tokenString string, err error) {
	authHeader := c.Request.Header.Get("Authorization")

	if checkIsStringEmpty(authHeader) {
		return tokenString, errors.New("Header authorization required")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return tokenString, errors.New("Header authorization format is invalid")
	}

	return parts[1], nil
}

func checkIsStringEmpty(input string) bool {
	return input == ""
}

func CheckJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := GetJwtTokenFromHeader(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &JwtData{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("JWT.SignatureKey")), nil
		})

		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		claims, ok := token.Claims.(*JwtData)
		if !ok || !token.Valid {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if claims.ExpiresAt == nil || claims.ExpiresAt.Unix() < time.Now().Unix() {
			c.JSON(http.StatusBadRequest, "Token expired, please login again")
			c.Abort()
			return
		}

		c.Set("Auth", claims)
		c.Next()
	}
}