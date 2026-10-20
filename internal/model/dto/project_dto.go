package dto

import (
	"mime/multipart"
	"time"

	"gorm.io/datatypes"
)

type ProjectResponse struct {
	UUID        string         `json:"uuid"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Image       string         `json:"image_url"`
	Link        string         `json:"link"`
	Repo        string         `json:"repo"`
	Category    datatypes.JSON `json:"category"`
	TechStack   datatypes.JSON `json:"techstack"`
	StartDate   time.Time      `json:"start_date"`
	EndDate     time.Time      `json:"end_date"`
	Company     string         `json:"company"`
}

type CreateProjectRequest struct {
	Title       string                `form:"title" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Image       *multipart.FileHeader `form:"image" binding:"required"`
	Link        string                `form:"live_url"`
	Repo        string                `form:"repo_url"`
	Category    string                `form:"category" binding:"required"`
	TechStack   string                `form:"techstack" binding:"required"`
	StartDate   time.Time             `form:"start_date" time_format:"2006-01-02" binding:"required"`
	EndDate     time.Time             `form:"end_date" time_format:"2006-01-02"`
	Company     string                `form:"company"`
}

type UpdateProjectRequest struct {
	UUID        string                `json:"-"`
	Title       string                `form:"title"`
	Description string                `form:"description"`
	Image       *multipart.FileHeader `form:"image"`
	Link        string                `form:"live_url"`
	Repo        string                `form:"repo_url"`
	Category    string                `form:"category"`
	TechStack   string                `form:"techstack"`
	StartDate   time.Time             `form:"start_date" time_format:"2006-01-02" `
	EndDate     time.Time             `form:"end_date" time_format:"2006-01-02"`
	Company     string                `form:"company"`
}

type UuidProjectRequest struct {
	UUID string `uri:"uuid" binding:"required"`
}