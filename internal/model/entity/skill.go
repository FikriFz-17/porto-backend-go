package entity

import "time"

type Skill struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID      string    `gorm:"type:uuid;default:gen_random_uuid()" json:"uuid"`	  
	Name      string    `gorm:"type:varchar" json:"name"`
	Icon      string    `gorm:"type:text" json:"icon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Skill) TableName() string {
	return "skills"
}