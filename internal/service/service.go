package service

import (
	"main.go/internal/model"
	"main.go/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Unit interface {
}

type Group interface {
}

type Space interface {
}

type Service struct {
	Authorization
	Unit
	Group
	Space
}

func NewServices(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		// Unit:          NewUnitService(repo),
		// Group:         NewGroupService(repo),
		// Space:         NewSpaceService(repo),
	}
}
