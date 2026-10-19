package repository

import (
	"porto-backend-go/internal/model/entity"

	"gorm.io/gorm"
)
type SkillRepository interface {
	FindAll() ([]entity.Skill, error)
	Create(skill *entity.Skill) error
	
}

type skillRepository struct {
	db *gorm.DB
}

func NewSkillRepository(db *gorm.DB) SkillRepository {
	return &skillRepository{db: db}
}

func (r *skillRepository) Create(skill *entity.Skill) error {
	return r.db.Create(skill).Error
}

func (r *skillRepository) FindAll() ([]entity.Skill, error) {
	var skills []entity.Skill
	err := r.db.Find(&skills).Error
	if err != nil {
		return nil, err
	}
	return skills, nil
}