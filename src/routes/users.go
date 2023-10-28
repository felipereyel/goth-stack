package routes

import (
	"encoding/base64"
	"goth/src/controllers"
	"goth/src/models"

	"github.com/gofiber/fiber/v2"
)

const cookieName = "goth:jwt"

func setUser(c *fiber.Ctx, user models.User) {
	c.Locals("user", user)
}

func getUser(c *fiber.Ctx) (models.User, error) {
	user, ok := c.Locals("user").(models.User)
	if !ok {
		return models.EmptyUser, fiber.ErrUnauthorized
	}

	return user, nil
}

func loginHandler(uc *controllers.UserController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Redirect(uc.GetAuthorizeURL(""))
	}
}

func authRedirect(uc *controllers.UserController) fiber.Handler {
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

func redirectToAuth(uc *controllers.UserController, c *fiber.Ctx) error {
	c.ClearCookie(cookieName)
	b64State := base64.StdEncoding.EncodeToString([]byte(c.Path()))
	return c.Redirect(uc.GetAuthorizeURL(b64State))
}

func verifyAuth(uc *controllers.UserController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwt := c.Cookies(cookieName)
		if jwt == "" {
			return redirectToAuth(uc, c)
		}

		user, err := uc.VerifyCookie(jwt)
		if err != nil {
			return redirectToAuth(uc, c)
		}

		setUser(c, user)
		return c.Next()
	}
}
