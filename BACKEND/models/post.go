package models

import "time"

type Post struct {
	ID           uint       `json:"id"`
	UserID       uint       `json:"user_id"`
	Companies    string     `json:"companies"`
	Post_Title   string     `json:"post_title"`
	Post_Text    string     `json:"post_text"`
	CategoriesID uint       `json:"categories_id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	User         User       `json:"-"`
	Categories   Categories `json:"-"`
	Comment      []Comment  `json:"-"`
	Votes        []Votes    `json:"-"`
}
