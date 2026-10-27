package repository

import (
	"porto-backend-go/internal/model/entity"
	"porto-backend-go/pkg/pagination"

	"gorm.io/gorm"
)

type ExperienceRepository interface {
	FindAll(param pagination.QueryParams) ([]entity.Experience, int64, error)
	Create(experience *entity.Experience) error
	Delete(uuid string) error
	GetExperienceByUUID(uuid string) (*entity.Experience, error)
	Update(uuid string, experience *entity.Experience) error
}

type experienceRepository struct {
	db *gorm.DB
}

func NewExperienceRepository(db *gorm.DB) ExperienceRepository {
	return &experienceRepository{db: db}
}

func (r *experienceRepository) FindAll(param pagination.QueryParams) ([]entity.Experience, int64, error) {
	var experiences []entity.Experience
	var total int64
	query := r.db.Model(&entity.Experience{})
	if param.Search != "" {
		searchTerm := "%" + param.Search + "%"
		query = query.Where("job ILIKE ? OR company ILIKE ?", searchTerm, searchTerm)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := query.Order("start_date DESC").
		Limit(param.GetLimit()).
		Offset(param.GetOffset()).
		Find(&experiences).Error
	if err != nil {
		return nil, 0, err
	}
	return experiences, total, nil
}

func (r *experienceRepository) Create(experience *entity.Experience) error {
	return r.db.Create(experience).Error
}

func (r *experienceRepository) Delete(uuid string) error {
	return r.db.Where("uuid = ?", uuid).Delete(&entity.Experience{}).Error
}

func (r *experienceRepository) GetExperienceByUUID(uuid string) (*entity.Experience, error) {
	var experience entity.Experience
	err := r.db.Where("uuid = ?", uuid).First(&experience).Error
	if err != nil {
		return nil, err
	}
	return &experience, nil
}

func (r *experienceRepository) Update(uuid string, experience *entity.Experience) error {
	return r.db.Where("uuid = ?", uuid).Updates(experience).Error
}