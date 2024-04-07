package model

type Article struct {
	BaseModel

	Title   string `json:"title"`
	Status  int    `json:"status"`
	Image   string `json:"image"`
	Summary string `json:"summary"`
	Content string `json:"content"`
}
