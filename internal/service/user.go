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

func (s *UserService) GetUserByEmail(email string) (user *model.User, err error) {
	user = &model.User{}
	err = s.DB.Where("email = ?", email).First(user).Error
	return user, err
}

func (s *UserService) CreateUser(user *model.User) (err error) {
	err = s.DB.Create(user).Error
	return err
}

func (s *UserService) ListUsers() (users []*model.User, err error) {
	err = s.DB.Preload("Todos").Find(&users).Error
	return users, err
}
