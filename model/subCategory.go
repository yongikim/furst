package model

import (
	"gorm.io/gorm"
)

type SubCategory struct {
	gorm.Model
	Name           string `binding:"required" gorm:"unique;not null"`
	UserID         uint
	MainCategoryID uint
	Pays           []Pay
}
