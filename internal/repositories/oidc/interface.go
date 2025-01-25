package oidc

import (
	"time"
)

type UserInfoResponse struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type OIDC interface {
	GetAuthorizeURL(scope, state string) string
	GetLogoutURL() string
	GetUser(code string) (UserInfoResponse, time.Time, error)
}
