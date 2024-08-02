package service

import (
	"github.com/evertonbzr/api-golang/internal/model"
	"gorm.io/gorm"
)

type TodoService struct {
	DB *gorm.DB
}

func NewTodoService(db *gorm.DB) *TodoService {
	return &TodoService{
		DB: db,
	}
}

func (s *TodoService) Create(todo []model.Todo) error {
	return s.DB.Create(&todo).Error
}
