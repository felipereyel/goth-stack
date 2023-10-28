package routes

import (
	"encoding/base64"
	"goth/src/controllers"
	"goth/src/models"

	"github.com/gofiber/fiber/v2"
)

const cookieName = "goth:jwt"

func withAuth[C controllers.Controllers](uController *controllers.UserController, controller *C, handler func(*C, *fiber.Ctx, models.User) error) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		jwt := ctx.Cookies(cookieName)
		if jwt == "" {
			return redirectToAuth(uController, ctx, true)
		}

		user, err := uController.VerifyCookie(jwt)
		if err != nil {
			return redirectToAuth(uController, ctx, true)
		}

		return handler(controller, ctx, user)
	}
}

func redirectToAuth(uc *controllers.UserController, c *fiber.Ctx, saveState bool) error {
	var b64State string
	if saveState {
		b64State = base64.StdEncoding.EncodeToString([]byte(c.Path()))
	}

	c.ClearCookie(cookieName)
	return c.Redirect(uc.GetAuthorizeURL(b64State))
}

func loginHandler(uc *controllers.UserController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return redirectToAuth(uc, c, false)
	}
}

func redirectHandler(uc *controllers.UserController) fiber.Handler {
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
