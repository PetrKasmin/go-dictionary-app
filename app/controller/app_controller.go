package controller

import (
	"github.com/app-dictionary/app/repositories"
	"github.com/app-dictionary/pkg/database"
	"net/http"
)

type AppController struct {
	Meta             Meta
	Data             Data
	EmbedFS          http.FileSystem
	DictRepository   repositories.DictRepository
	LetterRepository repositories.LetterRepository
	WordRepository   repositories.WordRepository
	TagRepository    repositories.TagRepository
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
