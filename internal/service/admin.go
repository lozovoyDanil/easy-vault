package service

import (
	"main.go/internal/model"
	"main.go/internal/repository"
)

type AdminService struct {
	repository.Admin
	repository.Authorization
}

func NewAdminService(repo *repository.Repository) *AdminService {
	return &AdminService{
		Admin:         repo.Admin,
		Authorization: repo.Authorization,
	}
}

func (s *AdminService) AllUsers() ([]model.User, error) {
	return s.Admin.AllUsers()
}

func (s *AdminService) BanUser(id int) error {
	isBanned, err := s.Authorization.UserIsBanned(id)
	if err != nil {
		return err
	}
	if isBanned {
		return ErrUserAlreadyBanned
	}

	return s.Admin.BanUser(id)
}

func (s *AdminService) DeleteUser(id int) error {
	isBanned, err := s.Authorization.UserIsBanned(id)
	if err != nil {
		return err
	}
	if !isBanned {
		return ErrCannotDeleteUser
	}

	return s.Admin.DeleteUser(id)
}
