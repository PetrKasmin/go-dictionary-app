package middleware

import (
	"github.com/app-dictionary/app/helpers"
	"github.com/gofiber/fiber/v2"
)

func (m *AppMiddleware) StaticFiles(c *fiber.Ctx) error {
	data, err := helpers.GetStaticFiles(m.EmbedFS, c.Params("dictionary"))
	if err == nil && data != nil {
		return c.Send(data)
	}

	return c.Next()
}
