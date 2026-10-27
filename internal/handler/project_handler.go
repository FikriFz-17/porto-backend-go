package handler

import (
	"net/http"
	"porto-backend-go/internal/model/dto"
	"porto-backend-go/internal/service"
	"porto-backend-go/pkg/pagination"
	"porto-backend-go/pkg/response"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	service service.ProjectService
}

func NewProjectHandler(service service.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: service}
}

// GetAllProjects godoc
// @Summary      Ambil daftar projects
// @Tags         Projects
// @Produce      json
// @Param        page    query    int    false "Halaman"
// @Param        limit   query    int    false "Limit"
// @Param        search  query    string false "Cari judul/perusahaan project"
// @Success      200     {object} map[string]interface{}
// @Router       /projects [get]
func (h *ProjectHandler) GetAllProjects(c *gin.Context) {
	var param pagination.QueryParams
	if err := c.ShouldBindQuery(&param); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	projects, meta, err := h.service.GetAllProjects(param)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessWithMeta(c, projects, meta)
}

// CreateProject godoc
// @Summary      Tambah project baru
// @Tags         Projects
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        title       formData string true  "Judul project"
// @Param        description formData string true  "Deskripsi project"
// @Param        image       formData file   true  "File gambar cover project"
// @Param        live_url    formData string false "URL Live Demo"
// @Param        repo_url    formData string false "URL Repository Git"
// @Param        category    formData string true  "JSON array string kategori"
// @Param        techstack   formData string true  "JSON array string teknologi"
// @Param        start_date  formData string true  "Tanggal mulai (YYYY-MM-DD)"
// @Param        end_date    formData string false "Tanggal selesai (YYYY-MM-DD)"
// @Param        company     formData string false "Nama perusahaan"
// @Success      200         {object} map[string]interface{}
// @Router       /projects [post]
func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var req dto.CreateProjectRequest
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.CreateProject(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, res)
}

// DeleteProject godoc
// @Summary      Hapus project
// @Tags         Projects
// @Security     BearerAuth
// @Param        uuid path     string true "UUID Project"
// @Success      200  {object} map[string]interface{}
// @Router       /projects/{uuid} [delete]
func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	var req dto.UuidProjectRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.DeleteProject(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Project deleted successfully")
}

// GetProjectByUUID godoc
// @Summary      Detail project
// @Tags         Projects
// @Param        uuid path     string true "UUID Project"
// @Success      200  {object} map[string]interface{}
// @Router       /projects/{uuid} [get]
func (h *ProjectHandler) GetProjectByUUID(c *gin.Context) {
	var req dto.UuidProjectRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.GetProjectByUUID(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, res)
}

// UpdateProject godoc
// @Summary      Perbarui project
// @Tags         Projects
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Param        uuid        path     string true  "UUID Project"
// @Param        title       formData string false "Judul project"
// @Param        description formData string false "Deskripsi project"
// @Param        image       formData file   false "File gambar baru"
// @Param        live_url    formData string false "URL Live Demo"
// @Param        repo_url    formData string false "URL Repository Git"
// @Param        category    formData string false "JSON array string kategori"
// @Param        techstack   formData string false "JSON array string teknologi"
// @Param        start_date  formData string false "Tanggal mulai (YYYY-MM-DD)"
// @Param        end_date    formData string false "Tanggal selesai (YYYY-MM-DD)"
// @Param        company     formData string false "Nama perusahaan"
// @Success      200         {object} map[string]interface{}
// @Router       /projects/{uuid} [put]
func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		response.Error(c, http.StatusBadRequest, "UUID is required")
		return
	}
	var req dto.UpdateProjectRequest
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	req.UUID = uuid
	res, err := h.service.UpdateProject(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, res)
}

// GetLatestProjects godoc
// @Summary      Ambil 6 project terbaru
// @Tags         Projects
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Router       /projects/latest [get]
func (h *ProjectHandler) GetLatestProjects(c *gin.Context) {
	projects, err := h.service.GetLatestProjects()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, projects)
}
