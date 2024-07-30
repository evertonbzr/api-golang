package service

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type TodoService struct {
	DB    *gorm.DB
	Cache *redis.Client
}

func NewTodoService() *TodoService {
	return &TodoService{}
}
