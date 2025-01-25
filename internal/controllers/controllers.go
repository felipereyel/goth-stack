package controllers

type Controllers interface {
	UserController | TaskController
}
