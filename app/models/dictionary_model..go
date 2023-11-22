package models

type Dictionary struct {
	ID         int    `json:"id"`
	Slug       string `gorm:"type:varchar(255)" json:"slug"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	CountWords string `json:"count_words"`
	IsDivider  bool   `json:"is_divider"`
}
