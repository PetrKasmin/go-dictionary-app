package models

type Dictionary struct {
	//gorm.Model
	ID         int    `json:"id"`
	Slug       string `gorm:"type:varchar(255)" json:"slug"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	CountWords string `json:"count_words"`
}
