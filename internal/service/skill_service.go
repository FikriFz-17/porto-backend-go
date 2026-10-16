package service

import (
	"porto-backend-go/internal/model/dto"
	"porto-backend-go/internal/repository"
)

type SkillService interface {
	GetAllSkills() ([]dto.SkillResponse, error)
}

type skillService struct {
	repo repository.SkillRepository
}

func NewSkillService(repo repository.SkillRepository) SkillService {
	return &skillService{repo: repo}
}

func (s *skillService) GetAllSkills() ([]dto.SkillResponse, error) {
	skills, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var res []dto.SkillResponse
	for _, skill := range skills {
		res = append(res, dto.SkillResponse{
			Name:    skill.Name,
			IconURL: skill.Icon,
		})
	}

	return res, nil
}