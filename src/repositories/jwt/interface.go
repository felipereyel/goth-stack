package jwt

import (
	"time"
)

// type JWTClaims struct {
// 	Id string `json:"id"`
// 	gojwt.RegisteredClaims
// }

type JWT interface {
	GenerateJWT(sub string, expiration time.Time) (string, error)
	ParseJWT(token string) (string, error)
}
