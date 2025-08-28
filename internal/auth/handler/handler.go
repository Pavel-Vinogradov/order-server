package handler

import (
	"net/http"
	"order-server/internal/auth/dto"
	"order-server/internal/auth/service"
	"order-server/pkg/jwt"
	"order-server/pkg/request"
	"order-server/pkg/response"
	"strconv"
)

type Handler struct {
	service *service.AuthService
}

// Auth godoc
// @Summary Авторизация
// @Description Авторизация
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   product body dto.AuthRequest
// @Success 201 {object} dto.AuthResponse
// @Failure 400 {object} response.ErrorsResponse
// @Router /auth [post]
func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
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
	authResponse := &dto.AuthResponse{
		SessionID: auth.Session,
	}
	response.WriteJSON(w, http.StatusCreated, authResponse)
}

// Verify godoc
// @Summary Авторизация
// @Description Верификация
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   product body dto.AuthVerifyRequest
// @Success 201 {object} dto.AuthVerifyResponse
// @Failure 400 {object} response.ErrorsResponse
// @Router /auth/verify [post]
func (h *Handler) Verify(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := request.HandleBody[dto.AuthVerifyRequest](w, r)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	_, err = strconv.Atoi(body.Code)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	var getUser, errUserBySession = h.service.GetUserBySession(body.SessionID)
	if errUserBySession != nil {
		response.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	generateJWT, errGenerateJWT := jwt.GenerateJWT(getUser.UserId)
	if errGenerateJWT != nil {
		response.WriteError(w, http.StatusInternalServerError, "ошибка генерации токена")
		return
	}
	authResponse := &dto.AuthVerifyResponse{
		Token: generateJWT,
	}
	response.WriteJSON(w, http.StatusOK, authResponse)
}

func NewRouteAuth(service *service.AuthService, router *http.ServeMux) {
	handler := &Handler{service: service}

	router.HandleFunc("/auth", handler.Auth)
	router.HandleFunc("/auth/verify", handler.Verify)
}
