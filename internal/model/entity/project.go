package entity

import (
	"time"

	"gorm.io/datatypes"
)

type Project struct {
	ID        	uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID 		string 		   `gorm:"column:uuid;type:uuid;default:gen_random_uuid()" json:"uuid"`
	Title       string         `gorm:"type:varchar" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	Image       string         `gorm:"type:text" json:"image"`
	Link        string         `gorm:"type:varchar" json:"link_url"`
	Repo        string         `gorm:"type:varchar" json:"repo_url"`
	Category    datatypes.JSON `gorm:"type:json" json:"category"`
	TechStack   datatypes.JSON `gorm:"type:json" json:"techstack"`
	StartDate   time.Time      `gorm:"type:date" json:"start_date"`
	EndDate     time.Time      `gorm:"type:date" json:"end_date"`
	Company 	string 			`gorm:"type:varchar" json:"company"`
}