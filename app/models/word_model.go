package models

type Word struct {
	ID              int    `json:"id"`
	DictionaryID    int    `json:"dictionary_id"`
	PaginatorID     int    `json:"paginator_id"`
	Slug            string `gorm:"type:varchar(255)" json:"slug"`
	Title           string `json:"title"`
	Content         string `json:"content"`
	LetterID        int    `json:"letter_id"`
	PrevSlug        string `json:"prev_slug"`
	PrevTitle       string `json:"prev_title"`
	NextSlug        string `json:"next_slug"`
	NextTitle       string `json:"next_title"`
	DictionaryTitle string `json:"dictionary_title"`
	DictionarySlug  string `json:"dictionary_slug"`
}
