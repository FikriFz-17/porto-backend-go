package entity

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID      string    `gorm:"type:uuid;default:gen_random_uuid()" json:"uuid"`
	Username  string    `gorm:"type:varchar" json:"username"`
	Password  string    `gorm:"type:varchar" json:"password"`
	IsActive  bool      `gorm:"type:boolean" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}