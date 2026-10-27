package repository

import (
	"porto-backend-go/internal/model/entity"

	"porto-backend-go/pkg/pagination"

	"gorm.io/gorm"
)
type SkillRepository interface {
	FindAll(param pagination.QueryParams) ([]entity.Skill, int64, error)
	Create(skill *entity.Skill) error
	Delete(uuid string) error
	GetSkillByUUID(uuid string) (*entity.Skill, error)
	Update(uuid string, name string, icon string) error
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

func (r *skillRepository) FindAll(param pagination.QueryParams) ([]entity.Skill, int64, error) {
	var skills []entity.Skill
	var total int64
	query := r.db.Model(&entity.Skill{})
	if param.Search != "" {
		searchTerm := "%" + param.Search + "%"
		query = query.Where("name ILIKE ?", searchTerm)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := query.Order("created_at DESC").
		Limit(param.GetLimit()).
		Offset(param.GetOffset()).
		Find(&skills).Error
	if err != nil {
		return nil, 0, err
	}
	return skills, total, nil
}

func (r *skillRepository) Delete(uuid string) error {
	return r.db.Where("uuid = ?", uuid).Delete(&entity.Skill{}).Error
}

func (r *skillRepository) GetSkillByUUID(uuid string) (*entity.Skill, error) {
	var skill entity.Skill
	err := r.db.Where("uuid = ?", uuid).First(&skill).Error
	if err != nil {
		return nil, err
	}
	return &skill, nil
}

func (r *skillRepository) Update(uuid string, name string, icon string) error {
	return r.db.Where("uuid = ?", uuid).Updates(entity.Skill{Name: name, Icon: icon}).Error
}