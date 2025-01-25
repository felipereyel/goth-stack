package jwt

import (
	"goth/internal/config"
	"time"

	gojwt "github.com/golang-jwt/jwt/v5"
)

type jwt struct {
	bsecret []byte
}

func NewJWTRepo(cfg config.ServerConfigs) JWT {
	return &jwt{
		bsecret: []byte(cfg.JwtSecret),
	}
}

func (j *jwt) GenerateJWT(sub string, expiration time.Time) (string, error) {
	claims := &gojwt.RegisteredClaims{
		Subject:   sub,
		IssuedAt:  gojwt.NewNumericDate(time.Now()),
		ExpiresAt: gojwt.NewNumericDate(expiration),
	}

	jwtoken := gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims)
	return jwtoken.SignedString(j.bsecret)
}

func (j *jwt) ParseJWT(token string) (string, error) {
	claims := gojwt.RegisteredClaims{}

	_, err := gojwt.ParseWithClaims(token, &claims, func(token *gojwt.Token) (interface{}, error) {
		return j.bsecret, nil
	})

	if err != nil {
		return "", err
	}

	if claims.Subject == "" {
		return "", gojwt.ErrInvalidKey
	}

	return claims.Subject, nil
}
