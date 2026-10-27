package handler

import (
	"net/http"
	"porto-backend-go/internal/model/dto"
	"porto-backend-go/internal/service"
	"porto-backend-go/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// Login godoc
// @Summary      Login Admin
// @Description  Autentikasi admin untuk mendapatkan JWT Bearer token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.LoginRequest true "Login Credentials"
// @Success      200     {object} map[string]interface{}
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.Login(req)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	response.Success(c, res)
}