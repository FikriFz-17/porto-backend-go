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

// GetAllSkills godoc
// @Summary      Ambil daftar skills
// @Tags         Skills
// @Produce      json
// @Param        page    query    int    false "Halaman"
// @Param        limit   query    int    false "Limit"
// @Param        search  query    string false "Cari nama skill"
// @Success      200     {object} map[string]interface{}
// @Router       /skills [get]
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

// CreateSkill godoc
// @Summary      Tambah skill baru
// @Tags         Skills
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        name formData string true "Nama skill"
// @Param        icon formData file   true "File icon skill"
// @Success      200  {object} map[string]interface{}
// @Router       /skills [post]
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

// DeleteSkill godoc
// @Summary      Hapus skill
// @Tags         Skills
// @Security     BearerAuth
// @Param        uuid path     string true "UUID Skill"
// @Success      200  {object} map[string]interface{}
// @Router       /skills/{uuid} [delete]
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

// GetSkillByUUID godoc
// @Summary      Detail skill
// @Tags         Skills
// @Param        uuid path     string true "UUID Skill"
// @Success      200  {object} map[string]interface{}
// @Router       /skills/{uuid} [get]
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

// UpdateSkill godoc
// @Summary      Perbarui skill
// @Tags         Skills
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Param        uuid path     string true  "UUID Skill"
// @Param        name formData string false "Nama skill"
// @Param        icon formData file   false "File icon baru"
// @Success      200  {object} map[string]interface{}
// @Router       /skills/{uuid} [put]
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
