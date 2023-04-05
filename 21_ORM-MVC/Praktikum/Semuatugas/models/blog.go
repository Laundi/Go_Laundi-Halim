package models

import "github.com/jinzu/gorm"

type Blog struct {
	gorm.model
	UserID	int 	`json:"user_id"` form:"user_id"
	Title 	string	`form:"title" json:"title"`
	Content	string	`form:"content" json: "content"`
	User	User	`json:"user"`
}
