package model

import (
  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  Name string `binding:"required" gorm:"unique;not null"`
  Password string `binding:"required"`
  Pays []Pay
  Wallets []Wallet
  MainCategories []MainCategory
  SubCategories []SubCategory
}
