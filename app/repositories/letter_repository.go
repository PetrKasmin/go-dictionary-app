package repositories

import (
	"github.com/app-dictionary/app/models"
	"gorm.io/gorm"
)

type LetterRepository interface {
	GetAllByDict(dictID int) ([]models.Letter, error)
	GetByDict(dictID int, title string) (*models.Letter, error)
}

type letterRepository struct {
	db *gorm.DB
}

func NewLetterRepository(db *gorm.DB) LetterRepository {
	return &letterRepository{db: db}
}

func (d *letterRepository) GetAllByDict(dictID int) ([]models.Letter, error) {
	var letters []models.Letter
	err := d.db.Find(&letters, "dictionary_id = ?", dictID).Error
	if err != nil {
		return nil, err
	}
	return letters, nil
}

func (d *letterRepository) GetByDict(dictID int, title string) (*models.Letter, error) {
	var letter models.Letter
	err := d.db.Find(&letter, "title = ? AND dictionary_id = ?", title, dictID).Error
	if err != nil {
		return nil, err
	}
	return &letter, nil
}
