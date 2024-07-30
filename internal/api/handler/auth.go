package handler

import (
	"net/http"

	"github.com/evertonbzr/api-golang/internal/api/utils"
	"github.com/evertonbzr/api-golang/internal/config"
	"github.com/evertonbzr/api-golang/internal/service"
	"gorm.io/gorm"
)

type AuthHandler struct {
	Service *service.UserService
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		Service: service.NewUserService(db),
	}
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// code here

		utils.RespondWithJSON(w, http.StatusOK, config.NAME)
	}
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// code here
		utils.RespondWithJSON(w, http.StatusOK, nil)

	}
}
