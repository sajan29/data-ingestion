package models

type Post struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	IngestedAt string `json:"ingested_at"`
	Source    string `json:"source"`
}