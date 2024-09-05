package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rs/xid"
)

type Claims struct {
	UserName string `json:"email"`
	jwt.StandardClaims
}

type JWTOutput struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

func GenerateToken() string {
	sessionToken := xid.New().String()
	return sessionToken
}
