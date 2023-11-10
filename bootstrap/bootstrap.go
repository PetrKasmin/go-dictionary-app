package bootstrap

import (
	"github.com/app-dictionary/pkg/database"
	"github.com/app-dictionary/pkg/env"
	"github.com/app-dictionary/pkg/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"net/http"
)

func NewApplication(embedFS http.FileSystem) *fiber.App {

	env.SetupEnvFile(embedFS)
	database.SetupDatabase()

	engine := html.NewFileSystem(embedFS, ".html")

	app := fiber.New(fiber.Config{
		Views:        engine,
		GETOnly:      true,
		UnescapePath: true,
		ViewsLayout:  "views/partials/default",
	})

	app.Use("/public", filesystem.New(filesystem.Config{
		Root:       embedFS,
		PathPrefix: "public",
		Browse:     true,
	}))

	app.Use(recover.New())
	app.Use(logger.New())
	//app.Get("/dashboard", monitor.New())

	router.InstallRouter(app, embedFS)

	return app
}
