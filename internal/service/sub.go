package service

import (
	"main.go/internal/model"
	"main.go/internal/repository"
)

type SubService struct {
	repo repository.Subscription
}

func NewSubService(repo repository.Subscription) *SubService {
	return &SubService{repo: repo}
}

func (s *SubService) SubByUserId(id int) (model.Subscription, error) {
	return s.repo.SubByUserId(id)
}

func (s *SubService) CreateSub(input model.Subscription) error {
	return s.repo.CreateSub(input)
}

func (s *SubService) UpdateSub(input model.Subscription) error {
	return s.repo.UpdateSub(input)
}
