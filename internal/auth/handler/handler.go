package handler

import (
	"net/http"
	"order-server/internal/auth/dto"
	"order-server/internal/auth/service"
	"order-server/pkg/request"
	"order-server/pkg/response"
)

type Handler struct {
	service *service.AuthService
}

// Auth godoc
// @Summary Создать продукт
// @Description Авторизация
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   product body dto.AuthRequest true "Продукт"
// @Success 201 {object} entity.Auth
// @Failure 400 {object} response.ErrorsResponse
// @Router /auth [post]
func (h Handler) Auth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body, err := request.HandleBody[dto.AuthRequest](w, r)
	if err != nil {
		return
	}
	auth, err := h.service.Auth(body)
	if err != nil {
		return
	}
	response.WriteJSON(w, http.StatusCreated, auth)
}

func NewRouteAuth(service *service.AuthService, router *http.ServeMux) {
	handler := &Handler{service: service}

	router.HandleFunc("/auth", handler.Auth)
}
