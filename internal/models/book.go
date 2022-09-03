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

func GetBookByID(ID int64) (Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", ID).Find(&book)
	return book, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}

func UpdateBook(ID int64, bookUpdate *Book) Book {
	book, _ := GetBookByID(ID)
	if bookUpdate.Name != "" {
		book.Name = bookUpdate.Name
	}
	if bookUpdate.Author != "" {
		book.Author = bookUpdate.Author
	}
	if bookUpdate.Publication != "" {
		book.Publication = bookUpdate.Publication
	}
	db.Save(&book)
	return book
}
