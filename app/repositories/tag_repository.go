package repositories

import (
	"github.com/app-dictionary/app/models"
	"gorm.io/gorm"
)

type TagRepository interface {
	GetAll() ([]models.Tag, error)
	GetBySlug(slug string) (*models.Tag, error)
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db: db}
}

func (t *tagRepository) GetAll() ([]models.Tag, error) {
	var tags []models.Tag
	if err := t.db.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t *tagRepository) GetBySlug(slug string) (*models.Tag, error) {
	var tag models.Tag
	if err := t.db.Find(&tag, "slug = ?", slug).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}
