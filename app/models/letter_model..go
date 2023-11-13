package models

type Letter struct {
	ID           int    `json:"id"`
	DictionaryID int    `gorm:"index:idx_dictionary" json:"dictionary_id"`
	Title        string `gorm:"type:varchar(255)" json:"title"`
	IsActive     bool   `json:"is_active"`
}
