package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetCookie(c *fiber.Ctx, name string, value string, expiration time.Time) {
	c.Cookie(buildCookie(name, value, expiration))
}

func ClearCookie(c *fiber.Ctx, name string) {
	// c.ClearCookie(name) does not work - https://github.com/gofiber/fiber/issues/1127
	c.Cookie(buildCookie(name, "", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)))
}

func buildCookie(name string, value string, expires time.Time) *fiber.Cookie {
	cookie := new(fiber.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = expires
	// cookie.HTTPOnly = true
	// cookie.Path = "/api/v1/auth/"
	// cookie.Domain = "example.com"
	return cookie
}
