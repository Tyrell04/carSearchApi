package common

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/marcleonschulz/carSearchApi/config"
	"github.com/marcleonschulz/carSearchApi/exception"
	"time"
)

func GenerateToken(username string, roles []map[string]interface{}, config config.Config) string {
	jwtSecret := config.Jwt.Secret

	claims := jwt.MapClaims{
		"username": username,
		"roles":    roles,
		"exp":      time.Now().Add(time.Minute * time.Duration(config.Jwt.AccessTokenExpireDuration)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(jwtSecret))
	exception.PanicLogging(err)

	return tokenSigned
}
