package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func InitModels(db *gorm.DB) {
	err := db.AutoMigrate(&Book{})
	if err != nil {
		panic(err)
	}
}
