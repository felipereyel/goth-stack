package routes

import (
	"encoding/base64"
	"goth/src/components"
	"goth/src/controllers"
	"goth/src/models"

	"github.com/gofiber/fiber/v2"
)

const cookieName = "goth:jwt"

func withAuth[C controllers.Controllers](uController *controllers.UserController, controller *C, handler func(*C, *fiber.Ctx, models.User) error) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		jwt := ctx.Cookies(cookieName)
		if jwt == "" {
			return redirectToAuthLogin(uController, ctx, true)
		}

		user, err := uController.VerifyCookie(jwt)
		if err != nil {
			return redirectToAuthLogin(uController, ctx, true)
		}

		return handler(controller, ctx, user)
	}
}

func redirectToAuthLogin(uc *controllers.UserController, c *fiber.Ctx, saveState bool) error {
	var b64State string
	if saveState {
		b64State = base64.StdEncoding.EncodeToString([]byte(c.Path()))
	}

	c.ClearCookie(cookieName)
	return c.Redirect(uc.GetAuthorizeURL(b64State))
}

func loginHandler(uc *controllers.UserController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return redirectToAuthLogin(uc, c, false)
	}
}

func postLoginHandler(uc *controllers.UserController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		code := c.Query("code")
		if code == "" {
			return c.Redirect("/")
		}

		cookie, err := uc.GetCookie(cookieName, code)
		if err != nil {
			return c.Redirect("/")
		}

		c.Cookie(cookie)

		state := c.Query("state")
		if state != "" {
			locationBytes, err := base64.StdEncoding.DecodeString(state)
			if err != nil {
				return c.Redirect("/")
			}

			return c.Redirect(string(locationBytes))
		}

		return c.Redirect("/")
	}
}

func logoutHandler(uc *controllers.UserController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.ClearCookie(cookieName)
		return c.Redirect(uc.GetLogoutURL())
	}
}

func postLogoutHandler(uc *controllers.UserController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.ClearCookie(cookieName)
		return sendPage(c, components.PostLogoutPage())
	}
}
