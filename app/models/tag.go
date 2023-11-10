package models

type Tag struct {
	//gorm.Model
	ID    int64  `json:"id"`
	Title string `json:"title"`
}
