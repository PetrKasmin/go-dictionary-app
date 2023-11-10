package models

type Letter struct {
	//gorm.Model
	ID           int    `json:"id"`
	DictionaryID int    `gorm:"index:idx_dictionary" json:"dictionary_id"`
	Title        string `gorm:"type:varchar(255)" json:"title"`
}
