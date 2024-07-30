package service

import (
	"github.com/evertonbzr/api-golang/internal/model"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		DB: db,
	}
}

func (s *UserService) GetUserById(id int) (user *model.User, err error) {
	user = &model.User{}
	err = s.DB.First(user, id).Error
	return user, err
}
