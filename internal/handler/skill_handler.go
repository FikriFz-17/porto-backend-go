package handler

import (
	"net/http"
	"porto-backend-go/internal/service"
	"porto-backend-go/pkg/response"

	"github.com/gin-gonic/gin"
)

type SkillHandler struct {
	service service.SkillService
}

func NewSkillHandler(service service.SkillService) *SkillHandler {
	return &SkillHandler{service: service}
}

func (h *SkillHandler) GetAllSkills(c *gin.Context) {
	skills, err := h.service.GetAllSkills()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, skills)
}