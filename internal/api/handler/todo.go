package handler

import "github.com/evertonbzr/api-golang/internal/service"

type TodoHandler struct {
	Service *service.TodoService
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{}
}
