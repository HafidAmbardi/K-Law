package models

import "time"

type Comment struct {
	ID           uint      `json:"id"`
	PostID       uint      `json:"post_id"`
	Comment_Text string    `json:"comment_text"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Post         Post      `json:"-"`
}
