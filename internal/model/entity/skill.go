package entity

import "time"

type Skill struct {
	ID        string    `gorm:"primaryKey;type:varchar" json:"id"`
	Name      string    `gorm:"type:varchar" json:"name"`
	Icon      string    `gorm:"type:text" json:"icon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Skill) TableName() string {
	return "skills"
}