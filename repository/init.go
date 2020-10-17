package repository

import (
	"errors"
	"fmt"
	"furst/model"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
  err := godotenv.Load(".env")
  if err != nil {
    panic(err.Error())
  }

  dbName := os.Getenv("DB_NAME")
  dbUser := os.Getenv("DB_USER")
  dbPW := os.Getenv("DB_PW")
  dbHost := os.Getenv("DB_HOST")
  dsn := dbUser + ":" + dbPW + "@(" + dbHost + ")" + "/" + dbName + "?charset=utf8&parseTime=true"
  err = errors.New("")
  db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil && err.Error() != "" {
    log.Fatal(err.Error())
  }
  db.AutoMigrate(&model.Book{})
  db.AutoMigrate(&model.Pay{})
  db.AutoMigrate(&model.User{})
  db.AutoMigrate(&model.Wallet{})
  db.AutoMigrate(&model.MainCategory{})
  db.AutoMigrate(&model.SubCategory{})
  fmt.Println("init data base ok")
}
