package dto

import "time"


type ExperienceResponse struct {
	UUID        string    `json:"uuid"`
	Job         string    `json:"job"`
	Company     string    `json:"company"`
	StartDate   time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
}

type CreateExperienceRequest struct {
	Job       string     `json:"job" form:"job" binding:"required"`
	Company   string     `json:"company" form:"company" binding:"required"`
	StartDate time.Time  `json:"start_date" form:"start_date" time_format:"2006-01-02" binding:"required"`
	EndDate   *time.Time `json:"end_date" form:"end_date" time_format:"2006-01-02"`
}

type UpdateExperienceRequest struct {
	UUID      string     `json:"-"`
	Job       string     `json:"job" form:"job"`
	Company   string     `json:"company" form:"company"`
	StartDate time.Time  `json:"start_date" form:"start_date" time_format:"2006-01-02"`
	EndDate   *time.Time `json:"end_date" form:"end_date" time_format:"2006-01-02"`
}


type UuidExperienceRequest struct {
	UUID string `uri:"uuid" binding:"required"`
}

