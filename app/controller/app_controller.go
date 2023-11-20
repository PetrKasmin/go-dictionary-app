package controller

import (
	"github.com/app-dictionary/app/repositories"
	"github.com/app-dictionary/pkg/database"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type AppController struct {
	Meta             Meta
	Data             Data
	Error            Error
	EmbedFS          http.FileSystem
	DictRepository   repositories.DictRepository
	LetterRepository repositories.LetterRepository
	WordRepository   repositories.WordRepository
	TagRepository    repositories.TagRepository
}

func (ctr *AppController) Response() fiber.Map {
	return fiber.Map{
		"Meta": ctr.Meta,
		"Data": ctr.Data,
	}
}

func (ctr *AppController) ErrorResponse(statuses ...int) fiber.Map {
	if len(statuses) > 0 {
		ctr.Error.Status = statuses[0]
	} else {
		ctr.Error.Status = fiber.StatusInternalServerError
	}

	return fiber.Map{
		"Meta":  ctr.Meta,
		"Error": ctr.Error,
	}
}

func NewAppController(embedFS http.FileSystem) *AppController {
	return &AppController{
		Meta:             NewMeta(),
		Data:             Data{},
		EmbedFS:          embedFS,
		DictRepository:   repositories.NewDictRepository(database.DB),
		LetterRepository: repositories.NewLetterRepository(database.DB),
		WordRepository:   repositories.NewWordRepository(database.DB),
		TagRepository:    repositories.NewTagRepository(database.DB),
	}
}
