package services

import (
	"github.com/app-dictionary/app/models"
	"github.com/app-dictionary/pkg/database"
)

type WordResult struct {
	ID              int    `json:"id"`
	DictionaryID    int    `json:"dictionary_id"`
	LetterID        int    `json:"letter_id"`
	Slug            string `json:"slug"`
	Title           string `json:"title"`
	Content         string `json:"content"`
	PrevSlug        string `json:"prev_slug"`
	PrevTitle       string `json:"prev_title"`
	NextSlug        string `json:"next_slug"`
	NextTitle       string `json:"next_title"`
	DictionaryTitle string `json:"dictionary_title"`
	DictionarySlug  string `json:"dictionary_slug"`
}

type Link struct {
	Title          string `json:"title"`
	Slug           string `json:"slug"`
	DictionarySlug string `json:"dictionary_slug"`
}

func GetWordsByTitle(dictionaryId int, title string) []WordResult {
	var words []WordResult
	database.DB.Model(&models.Word{}).
		Distinct("words.id").
		Select("words.id, words.dictionary_id, words.letter_id, words.slug, words.title, words.content, dictionaries.title as dictionary_title, dictionaries.slug as dictionary_slug").
		Joins("left join dictionaries on dictionaries.id = words.dictionary_id").
		Where("words.title = ? AND words.dictionary_id <> ?", title, dictionaryId).
		Scan(&words)
	return words
}

func GetWordsSearch(title string) []WordResult {
	var words []WordResult
	database.DB.Model(&models.Word{}).
		Distinct("words.id").
		Select("words.id, words.dictionary_id, words.letter_id, words.slug, words.title, words.content, dictionaries.title as dictionary_title, dictionaries.slug as dictionary_slug").
		Joins("left join dictionaries on dictionaries.id = words.dictionary_id").
		Where("words.title LIKE ?", "%"+title+"%").
		Scan(&words)
	return words
}

func GetWord(dictionaryId int, title string) WordResult {
	var word WordResult
	database.DB.Model(&models.Word{}).
		Select("words.id, words.dictionary_id, words.letter_id, words.slug, words.title, words.content, dictionaries.title as dictionary_title, dictionaries.slug as dictionary_slug").
		Joins("left join dictionaries on dictionaries.id = words.dictionary_id").
		Where("words.slug = ? AND words.dictionary_id = ?", title, dictionaryId).
		Scan(&word)
	return word
}

func GetWordsByDictionaryID(dictionaryID int, letterID int, limit, offset int) []models.Word {
	var words []models.Word
	database.DB.Limit(limit).
		Offset(offset).
		Find(&words, "letter_id = ? AND dictionary_id = ?", letterID, dictionaryID)
	return words
}

func GetWordsByOnlyDictionaryID(dictionaryID int) []models.Word {
	var words []models.Word
	database.DB.Find(&words, "dictionary_id = ?", dictionaryID)
	return words
}

func GetWordNav(id int) (prev Link, next Link) {
	var words []WordResult
	database.DB.Model(&models.Word{}).
		Select("words.id, words.dictionary_id, words.letter_id, words.slug, words.title, words.content, dictionaries.title as dictionary_title, dictionaries.slug as dictionary_slug").
		Joins("left join dictionaries on dictionaries.id = words.dictionary_id").
		Where("words.id IN (?, ?)", id-1, id+1).
		Scan(&words)

	for _, w := range words {
		if w.ID < id {
			prev.Title = w.Title
			prev.Slug = w.Slug
			prev.DictionarySlug = w.DictionarySlug
		} else if w.ID > id {
			next.Title = w.Title
			next.Slug = w.Slug
			next.DictionarySlug = w.DictionarySlug
		}
	}
	return prev, next
}
