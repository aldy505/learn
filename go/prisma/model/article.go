package model

type Article struct {
	AuthorID   string `json:"authorID"`
	Title      string `json:"title"`
	Subheading string `json:"subheading"`
	Content    string `json:"content"`
}
