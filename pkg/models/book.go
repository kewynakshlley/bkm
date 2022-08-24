package models

import (
	"books/pkg/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (book *Book) CreateBook() *Book {
	db.NewRecord(book)
	db.Create(&book)
	return book
}

func GetBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetOneBook(bookId int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", bookId).Find(&book)
	return &book, db
}

func DeleteBook(bookId int64) Book {
	var book Book
	db.Where("ID=?", bookId).Delete(book)
	return book
}
