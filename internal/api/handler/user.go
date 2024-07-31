package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/evertonbzr/api-golang/internal/service"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		Service: service.NewUserService(db),
	}
}

func (h *UserHandler) GetMe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// userId := r.Context().Value("userId").(int)

		// user, err := h.Service.GetUserById(userId)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }

		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)

		// if err := json.NewEncoder(w).Encode(user); err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// }
	}
}

func (h *UserHandler) GetUserById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		user, err := h.Service.GetUserById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
