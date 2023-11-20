package controller

type Error struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
