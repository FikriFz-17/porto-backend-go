package repository

import (
	"porto-backend-go/internal/model/entity"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	FindAll() ([]entity.Project, error)
	Create(project *entity.Project) error
	Delete (uuid string) error
	GetProjectByUUID(uuid string) (*entity.Project, error)
	Update(uuid string, project *entity.Project) error
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) FindAll() ([]entity.Project, error) {
	var projects []entity.Project
	err := r.db.Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
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