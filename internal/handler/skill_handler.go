package handler

import (
	"net/http"
	"porto-backend-go/internal/model/dto"
	"porto-backend-go/internal/service"
	"porto-backend-go/pkg/pagination"
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
	var param pagination.QueryParams
	if err := c.ShouldBindQuery(&param); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	skills, meta, err := h.service.GetAllSkills(param)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessWithMeta(c, skills, meta)
}

func (h *SkillHandler) CreateSkill(c *gin.Context) {
	var req dto.CreateSkillRequest
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.CreateSkill(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, res)
}

func (h *SkillHandler) DeleteSkill(c *gin.Context) {
	var req dto.UuidSkillRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.DeleteSkill(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Skill deleted successfully")
}

func (h *SkillHandler) GetSkillByUUID(c *gin.Context) {
	var req dto.UuidSkillRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.GetSkillByUUID(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, res)
}

func (h *SkillHandler) UpdateSkill(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		response.Error(c, http.StatusBadRequest, "UUID is required")
		return
	}

	var req dto.UpdateSkillRequest
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	req.UUID = uuid

	res, err := h.service.UpdateSkill(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, res)
}

