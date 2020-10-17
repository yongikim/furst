package repository

import (
	"furst/model"
)

type BookRepository struct{}

func (BookRepository) SetBook(book *model.Book) error {
	result := db.Create(&book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (BookRepository) GetBookList() []model.Book {
	books := make([]model.Book, 0)
	result := db.Limit(10).Find(&books)
	if result.Error != nil {
		panic(result.Error)
	}
	return books
}

func (BookRepository) UpdateBook(newBook *model.Book) error {
	result := db.Save(&newBook)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (BookRepository) DeleteBook(id int) error {
	result := db.Delete(&model.Book{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
