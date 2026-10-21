package entity

import "time"

type Experience struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID      string    `gorm:"column:uuid;type:uuid;default:gen_random_uuid()" json:"uuid"`
	Job       string    `gorm:"column:job;type:varchar" json:"job"`
	Company   string    `gorm:"column:company;type:varchar" json:"company"`
	StartDate time.Time `gorm:"column:start_date;type:date" json:"start_date"`
	EndDate   time.Time `gorm:"column:end_date;type:date" json:"end_date"`
}