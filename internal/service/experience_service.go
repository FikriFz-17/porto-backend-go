package service

import (
	"errors"
	"porto-backend-go/internal/model/dto"
	"porto-backend-go/internal/model/entity"
	"porto-backend-go/internal/repository"
	"porto-backend-go/pkg/pagination"
)

type ExperienceService interface {
	GetAllExperience(param pagination.QueryParams) ([]dto.ExperienceResponse, pagination.Meta, error)
	CreateExperience(req dto.CreateExperienceRequest) (*dto.ExperienceResponse, error)
	DeleteExperience(req dto.UuidExperienceRequest) error
	GetExperienceByUUID(req dto.UuidExperienceRequest) (*dto.ExperienceResponse, error)
	UpdateExperience(req dto.UpdateExperienceRequest) (*dto.ExperienceResponse, error)
}

type experienceService struct {
	repo repository.ExperienceRepository
}

func NewExperienceService(repo repository.ExperienceRepository) ExperienceService {
	return &experienceService{repo: repo}
}

func (s *experienceService) GetAllExperience(param pagination.QueryParams) ([]dto.ExperienceResponse, pagination.Meta, error) {
	experiences, total, err := s.repo.FindAll(param)
	if err != nil {
		return nil, pagination.Meta{}, err
	}

	res := []dto.ExperienceResponse{}
	for _, experience := range experiences {
		res = append(res, dto.ExperienceResponse{
			UUID:        experience.UUID,
			Job:         experience.Job,
			Company:     experience.Company,
			StartDate:   experience.StartDate,
			EndDate:     experience.EndDate,
		})
	}
	meta := pagination.BuildMeta(total, param.GetPage(), param.GetLimit())
	return res, meta, nil
}

func (s *experienceService) CreateExperience(req dto.CreateExperienceRequest) (*dto.ExperienceResponse, error) {
	experience := entity.Experience{
		Job:         req.Job,
		Company:     req.Company,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
	}

	err := s.repo.Create(&experience)
	if err != nil {
		return nil, err
	}
	return &dto.ExperienceResponse{
		UUID:        experience.UUID,
		Job:         experience.Job,
		Company:     experience.Company,
		StartDate:   experience.StartDate,
		EndDate:     experience.EndDate,
	}, nil
}

func (s *experienceService) DeleteExperience(req dto.UuidExperienceRequest) error {
	experience, err := s.repo.GetExperienceByUUID(req.UUID)
	if err != nil {
		return err
	}
	if experience.UUID == "" {
		return errors.New("experience not found")
	}
	return s.repo.Delete(req.UUID)
}

func (s *experienceService) GetExperienceByUUID(req dto.UuidExperienceRequest) (*dto.ExperienceResponse, error) {
	experience, err := s.repo.GetExperienceByUUID(req.UUID)
	if err != nil {
		return nil, err
	}
	if experience.UUID == "" {
		return nil, errors.New("experience not found")
	}
	return &dto.ExperienceResponse{
		UUID:        experience.UUID,
		Job:         experience.Job,
		Company:     experience.Company,
		StartDate:   experience.StartDate,
		EndDate:     experience.EndDate,
	}, nil
}

func (s *experienceService) UpdateExperience(req dto.UpdateExperienceRequest) (*dto.ExperienceResponse, error) {
	existing, err := s.repo.GetExperienceByUUID(req.UUID)
	if err != nil {
		return nil, err
	}
		if existing.UUID == "" {
		return nil, errors.New("experience not found")
	}

	job := existing.Job
	if req.Job != "" {
		job = req.Job
	}
	company := existing.Company
	if req.Company != "" {
		company = req.Company
	}
	startDate := existing.StartDate
	if !req.StartDate.IsZero() {
		startDate = req.StartDate
	}
	endDate := existing.EndDate
	if req.EndDate != nil {
		endDate = req.EndDate
	}

	experience := entity.Experience{
		Job:       job,
		Company:   company,
		StartDate: startDate,
		EndDate:   endDate,
	}
	if err := s.repo.Update(req.UUID, &experience); err != nil {
		return nil, err
	}
	return &dto.ExperienceResponse{
		UUID:      existing.UUID,
		Job:       experience.Job,
		Company:   experience.Company,
		StartDate: experience.StartDate,
		EndDate:   experience.EndDate,
	}, nil

}