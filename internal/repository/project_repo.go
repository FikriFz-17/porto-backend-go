package repository

import (
	"porto-backend-go/internal/model/entity"

	"porto-backend-go/pkg/pagination"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	FindAll(param pagination.QueryParams) ([]entity.Project, int64, error)
	Create(project *entity.Project) error
	Delete(uuid string) error
	GetProjectByUUID(uuid string) (*entity.Project, error)
	Update(uuid string, project *entity.Project) error
	GetLatest() ([]entity.Project, error)
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) FindAll(param pagination.QueryParams) ([]entity.Project, int64, error) {
	var projects []entity.Project
	var total int64
	query := r.db.Model(&entity.Project{})
	if param.Search != "" {
		searchTerm := "%" + param.Search + "%"
		query = query.Where("title ILIKE ?", searchTerm)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := query.Order("created_at DESC").
		Limit(param.GetLimit()).
		Offset(param.GetOffset()).
		Find(&projects).Error
	if err != nil {
		return nil, 0, err
	}
	return projects, total, nil
}

func (r *projectRepository) Create(project *entity.Project) error {
	return r.db.Create(project).Error
}

func (r *projectRepository) Delete(uuid string) error {
	return r.db.Where("uuid = ?", uuid).Delete(&entity.Project{}).Error
}

func (r *projectRepository) GetProjectByUUID(uuid string) (*entity.Project, error) {
	var project entity.Project
	err := r.db.Where("uuid = ?", uuid).First(&project).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *projectRepository) Update(uuid string, project *entity.Project) error {
	return r.db.Where("uuid = ?", uuid).Updates(project).Error
}

func (r *projectRepository) GetLatest() ([]entity.Project, error) {
	var projects []entity.Project

	query := `
		SELECT 
			uuid, title, description, image, tech_stack, category, 
			repo_url, live_url, start_date, end_date
		FROM projects
		ORDER BY end_date DESC
		LIMIT 6
	`
	
	err := r.db.Raw(query).Scan(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}