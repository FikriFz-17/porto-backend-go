package handler

import (
	"net/http"
	"porto-backend-go/internal/model/dto"
	"porto-backend-go/internal/service"
	"porto-backend-go/pkg/pagination"
	"porto-backend-go/pkg/response"

	"github.com/gin-gonic/gin"
)

type ExperienceHandler struct {
	service service.ExperienceService
}

func NewExperienceHandler(service service.ExperienceService) *ExperienceHandler {
	return &ExperienceHandler{service: service}
}

// GetAllExperience godoc
// @Summary      Ambil daftar riwayat pengalaman kerja
// @Tags         Experiences
// @Produce      json
// @Param        page    query    int    false "Halaman"
// @Param        limit   query    int    false "Limit"
// @Param        search  query    string false "Cari posisi pekerjaan atau nama perusahaan"
// @Success      200     {object} map[string]interface{}
// @Router       /experiences [get]
func (h *ExperienceHandler) GetAllExperience(c *gin.Context) {
	var param pagination.QueryParams
	if err := c.ShouldBindQuery(&param); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	experiences, meta, err := h.service.GetAllExperience(param)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessWithMeta(c, experiences, meta)
}

// CreateExperience godoc
// @Summary      Tambah pengalaman kerja baru
// @Tags         Experiences
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        request body dto.CreateExperienceRequest true "Payload data pengalaman kerja"
// @Success      200     {object} map[string]interface{}
// @Router       /experiences [post]
func (h *ExperienceHandler) CreateExperience(c *gin.Context) {
	var req dto.CreateExperienceRequest
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.CreateExperience(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, res)
}

// DeleteExperience godoc
// @Summary      Hapus pengalaman kerja
// @Tags         Experiences
// @Security     BearerAuth
// @Param        uuid path     string true "UUID Experience"
// @Success      200  {object} map[string]interface{}
// @Router       /experiences/{uuid} [delete]
func (h *ExperienceHandler) DeleteExperience(c *gin.Context) {
	var req dto.UuidExperienceRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.DeleteExperience(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Experience deleted successfully")
}

// GetExperienceByUUID godoc
// @Summary      Detail pengalaman kerja
// @Tags         Experiences
// @Param        uuid path     string true "UUID Experience"
// @Success      200  {object} map[string]interface{}
// @Router       /experiences/{uuid} [get]
func (h *ExperienceHandler) GetExperienceByUUID(c *gin.Context) {
	var req dto.UuidExperienceRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.GetExperienceByUUID(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, res)
}

// UpdateExperience godoc
// @Summary      Perbarui pengalaman kerja
// @Tags         Experiences
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        uuid    path string true "UUID Experience"
// @Param        request body dto.UpdateExperienceRequest true "Payload data update pengalaman"
// @Success      200     {object} map[string]interface{}
// @Router       /experiences/{uuid} [put]
func (h *ExperienceHandler) UpdateExperience(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		response.Error(c, http.StatusBadRequest, "UUID is required")
		return
	}
	var req dto.UpdateExperienceRequest
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	req.UUID = uuid
	res, err := h.service.UpdateExperience(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, res)
}
