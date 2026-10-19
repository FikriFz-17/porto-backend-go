package handler

import (
	"net/http"
	"porto-backend-go/internal/service"
	"porto-backend-go/pkg/response"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	service service.ProjectService
}

func NewProjectHandler(service service.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: service}
}

func (h *ProjectHandler) GetAllProjects(c *gin.Context) {
	projects, err := h.service.GetAllProjects()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, projects)
}