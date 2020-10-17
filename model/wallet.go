package model

import (
  "gorm.io/gorm"
)

type Wallet struct {
  gorm.Model
  Name string `binding:"required" gorm:"unique:not nill"`
  Amount int64
  UserID uint
  Pays []Pay
}
