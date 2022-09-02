package models

import (
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func InitModels(database *gorm.DB) {
	db = database
	err := db.AutoMigrate(&Book{})
	if err != nil {
		panic(err)
	}
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
