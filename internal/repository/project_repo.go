package repository

import (
	"porto-backend-go/internal/model/entity"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	FindAll() ([]entity.Project, error)
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