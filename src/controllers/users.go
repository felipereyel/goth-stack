package controllers

import (
	"goth/src/models"
	"goth/src/repositories/database"
	"goth/src/repositories/jwt"
	"goth/src/repositories/oidc"

	"github.com/gofiber/fiber/v2"
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

func (uc *UserController) VerifyCookie(token string) (models.User, error) {
	id, err := uc.jwtRepo.ParseJWT(token)
	if err != nil {
		return models.EmptyUser, err
	}

	return uc.dbRepo.RetrieveUserById(id)
}

func (uc *UserController) GetCookie(cookieName, code string) (*fiber.Cookie, error) {
	userInfo, expiration, err := uc.oidcRepo.GetUser(code)
	if err != nil {
		return nil, err
	}

	user, err := uc.dbRepo.UpsertUser(userInfo.Email)
	if err != nil {
		return nil, err
	}

	token, err := uc.jwtRepo.GenerateJWT(user.ID, expiration)
	if err != nil {
		return nil, err
	}

	return &fiber.Cookie{
		Name:    cookieName,
		Value:   token,
		Expires: expiration,
	}, nil
}
