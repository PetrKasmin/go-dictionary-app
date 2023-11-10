package main

import (
	"embed"
	"fmt"
	"github.com/app-dictionary/bootstrap"
	"github.com/app-dictionary/pkg/env"
	"log"
	"net/http"
)

//go:embed public/* views/* .env
var content embed.FS

func main() {
	app := bootstrap.NewApplication(http.FS(content))
	log.Fatal(
		app.Listen(
			fmt.Sprintf(
				"%s:%s",
				env.GetEnv("APP_HOST", "127.0.0.1"),
				env.GetEnv("APP_PORT", "3000"),
			),
		),
	)
}
