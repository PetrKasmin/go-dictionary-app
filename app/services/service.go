package services

import (
	"github.com/app-dictionary/app/models"
	"github.com/app-dictionary/pkg/database"
)

func AllDictionaries() []models.Dictionary {
	var dictionaries []models.Dictionary
	database.DB.Find(&dictionaries)
	return dictionaries
}

func GetDictionary(slug string) models.Dictionary {
	var dictionaries models.Dictionary
	database.DB.Find(&dictionaries, "slug = ?", slug)
	return dictionaries
}

func GetDictionaryBySlug(slug string) models.Dictionary {
	var dictionaries models.Dictionary
	database.DB.Find(&dictionaries, "slug = ?", slug)
	return dictionaries
}

type Letter struct {
	ID           int    `json:"id"`
	DictionaryID int    `json:"dictionary_id"`
	Title        string `json:"title"`
	IsActive     bool   `json:"is_active"`
}

func GetDictionaryLetters(dictionaryID int) []Letter {
	var letters []Letter
	database.DB.Find(&letters, "dictionary_id = ?", dictionaryID)
	return letters
}

func GetChunks(page []models.Dictionary, size int) [][]models.Dictionary {
	var chunk [][]models.Dictionary
	chunkSize := (len(page) + size - 1) / size
	for i := 0; i < len(page); i += chunkSize {
		end := i + chunkSize
		if end > len(page) {
			end = len(page)
		}
		chunk = append(chunk, page[i:end])
	}

	return chunk
}

func GetLetter(dictionaryID int, title string) models.Letter {
	var letter models.Letter
	database.DB.Find(&letter, "title = ? AND dictionary_id = ?", title, dictionaryID)
	return letter
}
