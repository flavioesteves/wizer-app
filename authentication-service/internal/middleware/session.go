package middleware

import (
	"github.com/rs/xid"
)

func GenerateToken() string {
	sessionToken := xid.New().String()
	return sessionToken
}
