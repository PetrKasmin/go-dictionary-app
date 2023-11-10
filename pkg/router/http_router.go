// routes.go
package router

import (
	"github.com/app-dictionary/app/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"io"
	"net/http"
	"strings"
)

type HttpRouter struct {
	embedFS http.FileSystem
}

func (h HttpRouter) InstallRouter(app *fiber.App) {
	group := app.Group("", cors.New(), csrf.New())
	group.Get("/", controllers.RenderMain)
	group.Get("/tags", controllers.RenderTags)
	group.Get("/terms", controllers.RenderTags)
	group.Get("/:dictionary", controllers.RenderDictionary)
	group.Get("/:dictionary/:word", controllers.RenderWord)

	app.Get("/:file", func(c *fiber.Ctx) error {
		fileName := c.Params("file")

		if strings.HasPrefix(fileName, "public/") {
			// Открываем файл из папки public
			file, err := h.embedFS.Open(fileName)
			if err != nil {
				return err
			}
			defer file.Close()

			data, err := io.ReadAll(file)
			if err != nil {
				return err
			}

			// Отправляем []byte в ответе
			return c.Send(data)
		}

		return c.SendStatus(http.StatusNotFound)
	})
}

func NewHttpRouter(embedFS http.FileSystem) *HttpRouter {
	return &HttpRouter{embedFS}
}
