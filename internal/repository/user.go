package repository

import (
	"github.com/evertonbzr/api-golang/internal/model"
	"github.com/evertonbzr/api-golang/pkg/infra/db"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (s *UserRepository) GetUserById(id uint) (user *model.User, err error) {
	user = &model.User{}
	err = db.GetDB().First(user, id).Error
	return user, err
}

func (s *UserRepository) GetUserByEmail(email string) (user *model.User, err error) {
	user = &model.User{}
	err = db.GetDB().Where("email = ?", email).First(user).Error
	return user, err
}

func (s *UserRepository) CreateUser(user *model.User) (err error) {
	err = db.GetDB().Create(user).Error
	return err
}

func (s *UserRepository) ListUsers() (users []*model.User, err error) {
	err = db.GetDB().Find(&users).Error
	return users, err
}

func (s *UserRepository) ListNotAdminUsers() (users []*model.User, err error) {
	err = db.GetDB().Where("role = ?", "user").Find(&users).Error
	return
}
