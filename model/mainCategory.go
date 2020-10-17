package model

import (
  "gorm.io/gorm"
)

type MainCategory struct {
  gorm.Model
  Name string `binding:"required" gorm:"unique;not null"`
  UserID uint
  SubCategories []SubCategory
  Pays []Pay
}
