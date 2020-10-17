package model

import (
  "gorm.io/gorm"
  "time"
)

type Pay struct {
  gorm.Model
  Amount int64
  Date time.Time
  Content string
  Memo string
  CountIn bool
  Transfer bool
  Mf string
  UserID uint
  WalletID uint
  MainCategoryID uint
  SubCategoryID uint
}
