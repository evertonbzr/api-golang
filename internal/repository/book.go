package repository

import (
	"github.com/evertonbzr/api-golang/internal/model"
	"github.com/evertonbzr/api-golang/pkg/infra/db"
)

type BookRepository struct {
}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

func (s *BookRepository) Create(book []model.Book) error {
	return db.GetDB().Create(&book).Error
}

func (s *BookRepository) GetByID(id uint) (model.Book, error) {
	book := model.Book{}

	if err := db.GetDB().First(&book, id).Error; err != nil {
		return model.Book{
			ID: 0,
		}, err
	}

	return book, nil
}

func (s *BookRepository) Update(book model.Book) error {
	var bookDB model.Book

	if err := db.GetDB().First(&bookDB, book.ID).Error; err != nil {
		return err
	}

	return db.GetDB().Model(&bookDB).Updates(book).Error
}

func (s *BookRepository) List() ([]model.Book, error) {
	books := []model.Book{}

	if err := db.GetDB().Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}
