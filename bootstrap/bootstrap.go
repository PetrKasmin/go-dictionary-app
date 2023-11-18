package bootstrap

import (
	"github.com/app-dictionary/pkg/database"
	"github.com/app-dictionary/pkg/env"
	"github.com/app-dictionary/pkg/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"html/template"
	"net/http"
	"time"
)

func NewApplication(embedFS http.FileSystem) *fiber.App {

	env.SetupEnvFile(embedFS)
	database.SetupDatabase()

	engine := html.NewFileSystem(embedFS, ".gohtml")
	engine.AddFunc(
		"unescape", func(s string) template.HTML {
			return template.HTML(s)
		},
	)

	app := fiber.New(fiber.Config{
		Views:                 engine,
		GETOnly:               true,
		UnescapePath:          true,
		ViewsLayout:           "views/layouts/default",
		DisableStartupMessage: env.IsProduction(),
	})

	app.Use("/public", filesystem.New(filesystem.Config{
		Root:       embedFS,
		PathPrefix: "public",
		Browse:     true,
	}))

	app.Use(limiter.New(limiter.Config{
		Max:               5000,
		Expiration:        5 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	app.Use(cors.New())
	app.Use(csrf.New())
	app.Use(recover.New())

	if !env.IsProduction() {
		app.Use(logger.New())
	}

	router.InstallRouter(app, embedFS)

	return app
}
