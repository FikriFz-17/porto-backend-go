package service

import (
	"porto-backend-go/internal/model/dto"
	"porto-backend-go/internal/repository"
)

type ProjectService interface {
	GetAllProjects() ([]dto.ProjectResponse, error)
}

type projectService struct {
	repo repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) ProjectService {
	return &projectService{repo: repo}
}

func (s *projectService) GetAllProjects() ([]dto.ProjectResponse, error) {
	projects, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var res []dto.ProjectResponse
	for _, project := range projects {
		res = append(res, dto.ProjectResponse{
			UUID:        project.UUID,
			Title:       project.Title,
			Description: project.Description,
			Image:       project.Image,
			Link:        project.Link,
			Repo:        project.Repo,
			Category:    project.Category,
			TechStack:   project.TechStack,
			StartDate:   project.StartDate,
			EndDate:     project.EndDate,
			Company:     project.Company,
		})
	}

	return res, nil
}