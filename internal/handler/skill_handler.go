package handler

import (
	"net/http"
	"porto-backend-go/internal/model/dto"
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

func (h *SkillHandler) CreateSkill(c *gin.Context) {
	name := c.PostForm("name")
	if name == ""{
		response.Error(c, http.StatusBadRequest, "Required Name")
		return
	}

	file, err := c.FormFile("icon")
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Required Icon Image")
		return
	}

	req := dto.CreateSkillRequest{
		Name: name,
		IconFile: file,
	}

	res, err := h.service.CreateSkill(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, res)
}