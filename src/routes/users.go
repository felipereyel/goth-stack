package routes

import (
	"encoding/base64"
	"goth/src/controllers"
	"goth/src/models"

	"github.com/gofiber/fiber/v2"
)

const cookieName = "goth:jwt"

func withUser(handler func(user models.User, c *fiber.Ctx) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, ok := c.Locals("user").(models.User)
		if !ok {
			return fiber.ErrUnauthorized
		}

		return handler(user, c)
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

func withAuth(uc *controllers.UserController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwt := c.Cookies(cookieName)
		if jwt == "" {
			return redirectToAuth(uc, c, true)
		}

		user, err := uc.VerifyCookie(jwt)
		if err != nil {
			return redirectToAuth(uc, c, true)
		}

		c.Locals("user", user)
		return c.Next()
	}
}
