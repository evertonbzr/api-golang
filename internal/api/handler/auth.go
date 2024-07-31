package handler

import (
	"net/http"
	"time"

	"github.com/evertonbzr/api-golang/internal/api/types"
	"github.com/evertonbzr/api-golang/internal/api/utils"
	"github.com/evertonbzr/api-golang/internal/model"
	"github.com/evertonbzr/api-golang/internal/service"
	"github.com/golang-jwt/jwt"
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
		data := types.LoginRequest{}

		if err := utils.DecodeJSONBody(w, r, &data); err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, nil)
			return
		}

		user, err := h.Service.GetUserByEmail(data.Email)

		if err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, nil)
			return
		}

		if user.Password != data.Password {
			utils.RespondWithJSON(w, http.StatusUnauthorized,
				map[string]string{"error": "invalid credentials"})
			return
		}

		token, _ := service.NewAccessToken(service.UserClaims{
			Id: user.ID,
			StandardClaims: jwt.StandardClaims{
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			},
		})

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
	}
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := types.RegisterRequest{}

		if err := utils.DecodeJSONBody(w, r, &data); err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, nil)
			return
		}

		user, _ := h.Service.GetUserByEmail(data.Email)

		if user.ID != 0 {
			utils.RespondWithJSON(w, http.StatusConflict,
				map[string]string{"error": "user already exists"})
			return
		}

		user = &model.User{
			FullName: data.FullName,
			Email:    data.Email,
			Password: data.Password,
		}

		if err := h.Service.CreateUser(user); err != nil {
			utils.RespondWithJSON(w, http.StatusInternalServerError, nil)
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, nil)
	}
}
