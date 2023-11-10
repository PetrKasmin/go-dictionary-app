package router

import (
	"fmt"
	"github.com/app-dictionary/app/helpers"
	"github.com/app-dictionary/app/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type ApiRouter struct {
}

func (h ApiRouter) InstallRouter(app *fiber.App) {
	api := app.Group("/api", limiter.New())
	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello from api",
		})
	})

	api.Get("/search", func(ctx *fiber.Ctx) error {
		word := ctx.Params("word")

		fmt.Println("word", word)
		words := services.GetWordsSearch(word)

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"words": words,
		})
	})

	api.Get("/sitemap", func(ctx *fiber.Ctx) error {
		dictionaries := services.AllDictionaries()
		var urls []string
		for i, d := range dictionaries {
			if i > 200 {
				//break
			}
			urls = append(urls, fmt.Sprintf("/%s", d.Slug))
			words := services.GetWordsByOnlyDictionaryID(d.ID)
			for _, w := range words {
				urls = append(urls, fmt.Sprintf("/%s/%s", d.Slug, w.Slug))
			}

			//fmt.Println("sitemap for dict "+d.Title, d.ID, len(urls))
		}

		helpers.SiteMapGenerator(urls)

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "generate ok",
		})
	})
}

func NewApiRouter() *ApiRouter {
	return &ApiRouter{}
}
