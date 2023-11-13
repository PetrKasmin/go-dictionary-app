package middleware

import "github.com/gofiber/fiber/v2"

type Middleware interface {
	BasicAuth(c *fiber.Ctx) error
}
