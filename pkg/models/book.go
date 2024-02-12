package models

import (
	"bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type book struct {
	gorm.Model
	Name         string `gorm:"" json:"name"`
	Author       string `json:"author"`
	Publications string `json:"publications"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&book{})
}

func (b *book) CreateBook() *book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}
