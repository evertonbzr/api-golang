package handler

import "github.com/evertonbzr/api-golang/internal/service"

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
