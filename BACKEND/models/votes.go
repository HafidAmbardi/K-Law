package models

type Votes struct {
	ID        uint `json:"id"`
	PostID    uint `json:"post_id"`
	UserID    uint `json:"user_id"`
	Vote_Type bool `json:"vote_type"`
	Post      Post `json:"-"`
	User      User `json:"-"`
}
