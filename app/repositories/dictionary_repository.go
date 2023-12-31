package repositories

import (
	"github.com/app-dictionary/app/models"
	"gorm.io/gorm"
)

type DictRepository interface {
	GetBySlug(slug string) (*models.Dictionary, error)
	GetAll() ([]models.Dictionary, error)
	GetByTag(slug string) ([]models.Dictionary, error)
}

type dictRepository struct {
	db *gorm.DB
}

func NewDictRepository(db *gorm.DB) DictRepository {
	return &dictRepository{db: db}
}

func (d *dictRepository) GetAll() ([]models.Dictionary, error) {
	var dictionaries []models.Dictionary
	if err := d.db.Find(&dictionaries).Error; err != nil {
		return nil, err
	}
	return dictionaries, nil
}

func (d *dictRepository) GetBySlug(slug string) (*models.Dictionary, error) {
	var dictionary models.Dictionary
	if err := d.db.Find(&dictionary, "slug = ?", slug).Error; err != nil {
		return nil, err
	}
	return &dictionary, nil
}

func (d *dictRepository) GetByTag(slug string) ([]models.Dictionary, error) {
	var dictionaries []models.Dictionary
	err := d.db.Model(&dictionaries).
		Select("dictionaries.id, dictionaries.slug, dictionaries.title, dictionaries.author, dictionaries.count_words").
		Joins("left join dictionaries_tags on dictionaries_tags.dictionary_id = dictionaries.id").
		Joins("left join tags on tags.id = dictionaries_tags.tag_id").
		Where("tags.slug = ?", slug).
		Scan(&dictionaries).
		Error

	if err != nil {
		return nil, err
	}

	return dictionaries, nil
}
