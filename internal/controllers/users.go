package controllers

import (
	"errors"
	"goth/internal/models"
	"goth/internal/repositories/database"
	"goth/internal/repositories/jwt"
	"goth/internal/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	dbRepo  database.Database
	jwtRepo jwt.JWT
}

func NewUserController(dbRepo database.Database, jwtRepo jwt.JWT) *UserController {
	return &UserController{dbRepo, jwtRepo}
}

func (uc *UserController) VerifyJWTCookie(token string) (models.User, error) {
	id, err := uc.jwtRepo.ParseJWT(token)
	if err != nil {
		return models.EmptyUser, err
	}

	return uc.dbRepo.RetrieveUserById(id)
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (uc *UserController) Login(req UserRequest, cookieName string) (*fiber.Cookie, error) {
	user, err := uc.dbRepo.RetrieveUserByName(req.Username)
	if err != nil {
		return nil, err
	}

	// slow checking -> design feature of bcrypt
	if !utils.CheckPasswordHash(req.Password, user.PswdHash) {
		return nil, errors.New("bad username or password")
	}

	expiration := time.Now().Add(7 * 24 * time.Hour)
	jwt, err := uc.jwtRepo.GenerateJWT(user.ID, expiration)
	if err != nil {
		return nil, err
	}

	return &fiber.Cookie{
		Expires: expiration,
		Name:    cookieName,
		Value:   jwt,
	}, nil
}

func (uc *UserController) Register(req UserRequest, cookieName string) (*fiber.Cookie, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		ID:       models.GenerateId(),
		Username: req.Username,
		PswdHash: hashedPassword,
	}

	err = uc.dbRepo.InsertUser(user)
	if err != nil {
		return nil, err
	}

	expiration := time.Now().Add(7 * 24 * time.Hour)
	jwt, err := uc.jwtRepo.GenerateJWT(user.ID, expiration)
	if err != nil {
		return nil, err
	}

	return &fiber.Cookie{
		Expires: expiration,
		Name:    cookieName,
		Value:   jwt,
	}, nil
}
