package models

type Categories struct {
	ID              uint   `json:"id"`
	Categories_Text string `json:"categories_text"`
	Categories_Desc string `json:"categories_desc"`
	Post            []Post `json:"-"`
}
