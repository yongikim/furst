package model

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}
