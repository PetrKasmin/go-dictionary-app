package controller

import (
	"github.com/app-dictionary/app/models"
	"github.com/app-dictionary/app/repositories"
)

type Data struct {
	Title               string              `json:"title,omitempty"`
	Dictionaries        []models.Dictionary `json:"dictionaries,omitempty"`
	DictionariesLetters []string            `json:"dictionaries_letters,omitempty"`
	Tags                []models.Tag        `json:"tags,omitempty"`
	Dictionary          *models.Dictionary  `json:"dictionary,omitempty"`
	DictionaryWords     []models.Word       `json:"dictionary_words,omitempty"`
	Letters             []models.Letter     `json:"letters,omitempty"`
	Letter              string              `json:"letter,omitempty"`
	Words               []models.Word       `json:"words,omitempty"`
	Word                *models.Word        `json:"word,omitempty"`
	Page                int                 `json:"page,omitempty"`
	Prev                int                 `json:"prev,omitempty"`
	Next                int                 `json:"next,omitempty"`
	PrevLink            repositories.Link   `json:"prev_link,omitempty"`
	NextLink            repositories.Link   `json:"next_link,omitempty"`
	CanPrevPage         bool                `json:"can_prev_page,omitempty"`
	CanNextPage         bool                `json:"can_next_page,omitempty"`
}
