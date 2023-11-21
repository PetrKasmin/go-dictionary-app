package models

type Tag struct {
	ID    int64  `json:"id,omitempty"`
	Slug  string `json:"slug,omitempty"`
	Title string `json:"title,omitempty"`
}
