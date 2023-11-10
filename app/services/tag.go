package services

import (
	"github.com/app-dictionary/app/models"
	"github.com/app-dictionary/pkg/database"
)

func AllTags() []models.Tag {
	var tags []models.Tag
	database.DB.Find(&tags)
	return tags
}

func GetTag(id int64) models.Tag {
	var tag models.Tag
	database.DB.Find(&tag, id)
	return tag
}
