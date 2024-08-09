package service

import (
	"github.com/evertonbzr/api-golang/internal/model"
	"gorm.io/gorm"
)

type BookService struct {
	DB *gorm.DB
}

func NewBookService(db *gorm.DB) *BookService {
	return &BookService{
		DB: db,
	}
}

func (s *BookService) Create(book []model.Book) error {
	return s.DB.Create(&book).Error
}

func (s *BookService) GetByID(id uint) (model.Book, error) {
	book := model.Book{}

	if err := s.DB.First(&book, id).Error; err != nil {
		return model.Book{
			ID: 0,
		}, err
	}

	return book, nil
}

func (s *BookService) Update(book model.Book) error {
	return s.DB.Save(&book).Error
}

func (s *BookService) List() ([]model.Book, error) {
	books := []model.Book{}

	if err := s.DB.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}
