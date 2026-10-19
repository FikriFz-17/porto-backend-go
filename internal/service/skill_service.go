package service

import (
	"porto-backend-go/internal/model/dto"
	"porto-backend-go/internal/model/entity"
	"porto-backend-go/internal/repository"
	"porto-backend-go/pkg/storage"
)

type SkillService interface {
	GetAllSkills() ([]dto.SkillResponse, error)
	CreateSkill(req dto.CreateSkillRequest) (*dto.SkillResponse, error)
}

type skillService struct {
	repo repository.SkillRepository
	storage storage.StorageService
}

func NewSkillService(repo repository.SkillRepository, storage storage.StorageService) SkillService {
	return &skillService{repo: repo, storage: storage}
}

func (s *skillService) GetAllSkills() ([]dto.SkillResponse, error) {
	skills, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var res []dto.SkillResponse
	for _, skill := range skills {
		res = append(res, dto.SkillResponse{
			UUID:    skill.UUID,
			Name:    skill.Name,
			IconURL: skill.Icon,
		})
	}

	return res, nil
}

func (s *skillService) CreateSkill(req dto.CreateSkillRequest) (*dto.SkillResponse, error) {
	iconURL, err := s.storage.UploadFile(req.IconFile, "skills")
	if err != nil {
		return nil, err
	}

	skill := entity.Skill{
		Name: req.Name,
		Icon: iconURL,
	}

	err = s.repo.Create(&skill)

	if err != nil {
		return nil, err
	}

	return &dto.SkillResponse{
		UUID:    skill.UUID,
		Name:    skill.Name,
		IconURL: skill.Icon,
	}, nil
}