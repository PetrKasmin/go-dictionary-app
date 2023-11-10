package models

type Word struct {
	//gorm.Model gorm:"type:varchar(255)"
	ID           int    `json:"id"`
	DictionaryID int    `json:"dictionary_id"`
	PaginatorID  int    `json:"paginator_id"`
	Slug         string `gorm:"type:varchar(255)" json:"slug"`
	Title        string `json:"title"`
	Content      string `json:"content"`
}
