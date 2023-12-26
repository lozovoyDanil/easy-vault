package service

import (
	"time"

	"main.go/internal/model"
	"main.go/internal/repository"
)

type PartService struct {
	repo repository.Partnership
}

func NewPartService(repo repository.Partnership) *PartService {
	return &PartService{repo: repo}
}

func (s *PartService) PartByUserId(id int) (model.Partnership, error) {
	return s.repo.PartByUserId(id)
}

func (s *PartService) CreatePart(userId, tier int) error {
	userPart, err := s.PartByUserId(userId)
	if err != nil {
		return err
	}
	if userPart.UserId != 0 {
		return ErrPartnershipViolation
	}

	part := model.Partnership{
		UserId:    userId,
		Tier:      tier,
		StartedAt: time.Now(),
	}
	return s.repo.CreatePart(part)
}

func (s *PartService) UpdatePart(userId, tier int) error {
	userPart, err := s.PartByUserId(userId)
	if err != nil {
		return err
	}
	if userPart.Tier >= tier {
		return ErrTierViolation
	}

	part := model.Partnership{
		UserId:    userId,
		Tier:      tier,
		StartedAt: time.Now(),
	}
	return s.repo.UpdatePart(part)
}
