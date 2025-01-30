package routes

import (
	"goth/internal/components"
	"goth/internal/controllers"
	"goth/internal/models"
	"goth/internal/utils"

	"github.com/gofiber/fiber/v2"
)

const cookieName = "goth:jwt"

func withAuth[C controllers.Controllers](uController *controllers.UserController, controller *C, handler func(*C, *fiber.Ctx, models.User) error) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		jwt := ctx.Cookies(cookieName)
		if jwt == "" {
			return ctx.Redirect("/auth/login")
		}

		user, err := uController.VerifyJWTCookie(jwt)
		if err != nil {
			return ctx.Redirect("/auth/login")
		}

		return handler(controller, ctx, user)
	}
}

func getLogoutHandler(c *fiber.Ctx) error {
	utils.ClearCookie(c, cookieName)
	return sendPage(c, components.PostLogoutPage())
}

func getRegisterHandler(c *fiber.Ctx) error {
	utils.ClearCookie(c, cookieName)
	return sendPage(c, components.RegisterPage())
}

func getLoginHandler(c *fiber.Ctx) error {
	utils.ClearCookie(c, cookieName)
	return sendPage(c, components.LoginPage())
}

func postLoginHandler(uc *controllers.UserController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req controllers.UserRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}

		cookie, err := uc.Login(req, cookieName)
		if err != nil {
			return err
		}

		c.Cookie(cookie)
		c.Set("HX-Redirect", "/")
		return c.SendStatus(fiber.StatusOK)
	}
}

func postRegisterHandler(uc *controllers.UserController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req controllers.UserRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}

		cookie, err := uc.Register(req, cookieName)
		if err != nil {
			return err
		}

		c.Cookie(cookie)
		c.Set("HX-Redirect", "/")
		return c.SendStatus(fiber.StatusOK)
	}
}
