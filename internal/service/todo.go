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

func (s *TodoService) GetByID(id uint) (model.Todo, error) {
	todo := model.Todo{}

	if err := s.DB.First(&todo, id).Error; err != nil {
		return model.Todo{
			ID: 0,
		}, err
	}

	return todo, nil
}

func (s *TodoService) Update(todo model.Todo) error {
	return s.DB.Save(&todo).Error
}

func (s *TodoService) List() ([]model.Todo, error) {
	todos := []model.Todo{}

	if err := s.DB.Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}
