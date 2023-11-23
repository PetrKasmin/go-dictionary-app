package middleware

import "net/http"

type AppMiddleware struct {
	EmbedFS http.FileSystem
}

func NewMiddleware(embedFS http.FileSystem) *AppMiddleware {
	return &AppMiddleware{embedFS}
}
