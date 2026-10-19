package dto

import (
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