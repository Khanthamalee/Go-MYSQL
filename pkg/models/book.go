package model

import (
	"goandmysql/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	NAME        string `*gorm: "" json: "name`
	AUTHOR      string `json : "author"`
	PUBLICATION string `json : "publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
