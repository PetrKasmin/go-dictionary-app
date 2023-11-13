package middleware

import (
	"github.com/app-dictionary/pkg/env"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func (m *AppMiddleware) BasicAuth(c *fiber.Ctx) error {
	login := env.GetEnv("ADMIN_LOGIN", "")
	password := env.GetEnv("ADMIN_PASSWORD", "")

	if login == "" || password == "" {
		return c.SendStatus(fiber.StatusForbidden)
	}

	auth := basicauth.New(basicauth.Config{
		Users: map[string]string{
			login: password,
		},
	})

	return auth(c)
}
