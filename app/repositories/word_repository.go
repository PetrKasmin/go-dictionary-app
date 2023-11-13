package repositories

import (
	"github.com/app-dictionary/app/models"
	"gorm.io/gorm"
)

type WordRepository interface {
	GetByDictAndLetter(dictID, letterID, limit, offset int) ([]models.Word, error)
	GetWordsByDictAndTitle(dictID int, title string) ([]models.Word, error)
	GetWordByDictAndTitle(dictID int, slug string) (*models.Word, error)
	GetWordForNav(wordID int) ([]models.Word, error)
	GetByDict(dictID int) ([]models.Word, error)
	Search(title string) ([]models.Word, error)
}

type wordRepository struct {
	db *gorm.DB
}

type Link struct {
	Title          string `json:"title"`
	Slug           string `json:"slug"`
	DictionarySlug string `json:"dictionary_slug"`
}

func NewWordRepository(db *gorm.DB) WordRepository {
	return &wordRepository{db: db}
}

func (w *wordRepository) GetByDictAndLetter(dictID, letterID, limit, offset int) ([]models.Word, error) {
	var words []models.Word
	err := w.db.Limit(limit).
		Offset(offset).
		Find(&words, "letter_id = ? AND dictionary_id = ?", letterID, dictID).
		Error

	if err != nil {
		return nil, err
	}

	return words, nil
}

func (w *wordRepository) GetByDict(dictID int) ([]models.Word, error) {
	var words []models.Word
	if err := w.db.Find(&words, "dictionary_id = ?", dictID).Error; err != nil {
		return nil, err
	}
	return words, nil
}

func (w *wordRepository) GetWordByDictAndTitle(dictID int, slug string) (*models.Word, error) {
	var word models.Word
	err := w.db.Model(&word).
		Select("words.id, words.dictionary_id, words.letter_id, words.slug, words.title, words.content, dictionaries.title as dictionary_title, dictionaries.slug as dictionary_slug").
		Joins("left join dictionaries on dictionaries.id = words.dictionary_id").
		Where("words.slug = ? AND words.dictionary_id = ?", slug, dictID).
		Scan(&word).
		Error

	if err != nil {
		return nil, err
	}

	return &word, nil
}

func (w *wordRepository) GetWordsByDictAndTitle(dictId int, title string) ([]models.Word, error) {
	var words []models.Word
	err := w.db.Model(&words).
		Distinct("words.id").
		Select("words.id, words.dictionary_id, words.letter_id, words.slug, words.title, words.content, dictionaries.title as dictionary_title, dictionaries.slug as dictionary_slug").
		Joins("left join dictionaries on dictionaries.id = words.dictionary_id").
		Where("words.title = ? AND words.dictionary_id <> ?", title, dictId).
		Scan(&words).
		Error

	if err != nil {
		return nil, err
	}

	return words, nil
}

func (w *wordRepository) Search(title string) ([]models.Word, error) {
	var words []models.Word
	err := w.db.Model(&models.Word{}).
		Distinct("words.id").
		Select("words.id, words.dictionary_id, words.letter_id, words.slug, words.title, words.content, dictionaries.title as dictionary_title, dictionaries.slug as dictionary_slug").
		Joins("left join dictionaries on dictionaries.id = words.dictionary_id").
		Where("words.title LIKE ?", "%"+title+"%").
		Scan(&words).
		Error

	if err != nil {
		return nil, err
	}

	return words, nil
}

func (w *wordRepository) GetWordForNav(wordID int) ([]models.Word, error) {
	var words []models.Word
	err := w.db.Model(&words).
		Select("words.id, words.dictionary_id, words.letter_id, words.slug, words.title, words.content, dictionaries.title as dictionary_title, dictionaries.slug as dictionary_slug").
		Joins("left join dictionaries on dictionaries.id = words.dictionary_id").
		Where("words.id IN (?, ?)", wordID-1, wordID+1).
		Scan(&words).
		Error

	if err != nil {
		return nil, err
	}

	return words, nil
}
