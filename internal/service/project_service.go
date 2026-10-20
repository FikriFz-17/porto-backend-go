package service

import (
	"porto-backend-go/internal/model/dto"
	"porto-backend-go/internal/model/entity"
	"porto-backend-go/internal/repository"
	"porto-backend-go/pkg/storage"

	"gorm.io/datatypes"
)

type ProjectService interface {
	GetAllProjects() ([]dto.ProjectResponse, error)
	CreateProject(req dto.CreateProjectRequest) (*dto.ProjectResponse, error)
	DeleteProject(req dto.UuidProjectRequest) error
	GetProjectByUUID(req dto.UuidProjectRequest) (*dto.ProjectResponse, error)
	UpdateProject(req dto.UpdateProjectRequest) (*dto.ProjectResponse, error)
}

type projectService struct {
	repo repository.ProjectRepository
	storage storage.StorageService
}

func NewProjectService(repo repository.ProjectRepository, storage storage.StorageService) ProjectService {
	return &projectService{repo: repo, storage: storage}
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

func (s *projectService) CreateProject(req dto.CreateProjectRequest) (*dto.ProjectResponse, error) {
	imgURL, err := s.storage.UploadFile(req.Image, "projects")
	if err != nil {
		return nil, err
	}

	project := entity.Project {
		Title:       req.Title,
		Description: req.Description,
		Image:       imgURL,
		Link:        req.Link,
		Repo:        req.Repo,
		Category:    datatypes.JSON(req.Category),
		TechStack:   datatypes.JSON(req.TechStack),
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Company:     req.Company,
	}


	err = s.repo.Create(&project)
	if err != nil {
		return nil, err
	}

	return &dto.ProjectResponse{
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
	}, nil
}

func (s *projectService) DeleteProject(req dto.UuidProjectRequest) error {
	project, err := s.repo.GetProjectByUUID(req.UUID)
	if err != nil {
		return err
	}
	if project.Image != "" {
		_ = s.storage.DeleteFile(project.Image)
	}
	return s.repo.Delete(req.UUID)
}

func (s *projectService) GetProjectByUUID(req dto.UuidProjectRequest) (*dto.ProjectResponse, error) {
	project, err := s.repo.GetProjectByUUID(req.UUID)
	if err != nil {
		return nil, err
	}
	return &dto.ProjectResponse{
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
	}, nil
}

func (s *projectService) UpdateProject(req dto.UpdateProjectRequest) (*dto.ProjectResponse, error) {
	existing, err := s.repo.GetProjectByUUID(req.UUID)
	if err != nil {
		return nil, err
	}

	imageURL := existing.Image
	if req.Image != nil {
		newImageURL, err := s.storage.UploadFile(req.Image, "projects")
		if err != nil {
			return nil, err
		}
		if existing.Image != "" {
			_ = s.storage.DeleteFile(existing.Image)
		}
		imageURL = newImageURL
	}

	category := existing.Category
	if req.Category != "" {
		category = datatypes.JSON(req.Category)
	}
	techStack := existing.TechStack
	if req.TechStack != "" {
		techStack = datatypes.JSON(req.TechStack)
	}
	title := existing.Title
	if req.Title != "" {
		title = req.Title
	}
	description := existing.Description
	if req.Description != "" {
		description = req.Description
	}
	project := entity.Project{
		Title:       title,
		Description: description,
		Image:       imageURL,
		Link:        req.Link,
		Repo:        req.Repo,
		Category:    category,
		TechStack:   techStack,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Company:     req.Company,
	}

	if err := s.repo.Update(req.UUID, &project); err != nil {
		return nil, err
	}
	return &dto.ProjectResponse{
		UUID:        existing.UUID,
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
	}, nil
}