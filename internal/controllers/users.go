package controllers

import (
	"goth/internal/models"
	"goth/internal/repositories/database"
	"goth/internal/repositories/jwt"
	"goth/internal/repositories/oidc"
	"time"
)

type UserController struct {
	dbRepo   database.Database
	oidcRepo oidc.OIDC
	jwtRepo  jwt.JWT
}

func NewUserController(dbRepo database.Database, oidcRepo oidc.OIDC, jwtRepo jwt.JWT) *UserController {
	return &UserController{dbRepo, oidcRepo, jwtRepo}
}

func (uc *UserController) GetAuthorizeURL(b64State string) string {
	return uc.oidcRepo.GetAuthorizeURL("openid profile email", b64State)
}

func (uc *UserController) GetLogoutURL() string {
	return uc.oidcRepo.GetLogoutURL()
}

func (uc *UserController) VerifyJWTCookie(token string) (models.User, error) {
	id, err := uc.jwtRepo.ParseJWT(token)
	if err != nil {
		return models.EmptyUser, err
	}

	return uc.dbRepo.RetrieveUserById(id)
}

func (uc *UserController) GetJWTCookie(code string) (string, time.Time, error) {
	userInfo, expiration, err := uc.oidcRepo.GetUser(code)
	if err != nil {
		return "", time.Time{}, err
	}

	user, err := uc.dbRepo.UpsertUser(userInfo.Email)
	if err != nil {
		return "", time.Time{}, err
	}

	token, err := uc.jwtRepo.GenerateJWT(user.ID, expiration)
	if err != nil {
		return "", time.Time{}, err
	}

	return token, expiration, nil
}
