package service

import (
	"main.go/internal/model"
	"main.go/internal/repository"
)

type UnitService struct {
	repository.Unit
	//* Необходима для проверки принадлежности Group пользователю
	repository.Group
}

func NewUnitService(repo *repository.Repository) *UnitService {
	return &UnitService{
		Unit:  repo.Unit,
		Group: repo.Group,
	}
}

func (s *UnitService) GroupUnits(userId, groupId int) ([]model.StorageUnit, error) {
	if err := s.Group.GroupBelongsToUser(userId, groupId); err != nil {
		return nil, err
	}

	return s.Unit.GroupUnits(groupId)
}

func (s *UnitService) UnitById(userId, unitId int) (model.StorageUnit, error) {
	if err := s.Unit.UnitBelongsToUser(userId, unitId); err != nil {
		return model.StorageUnit{}, err
	}

	return s.Unit.UnitById(unitId)
}

func (s *UnitService) CreateUnit(userId, groupId int, unit model.StorageUnit) (int, error) {
	if err := s.Group.GroupBelongsToUser(userId, groupId); err != nil {
		return 0, err
	}

	unit.GroupId = groupId

	return s.Unit.CreateUnit(unit)
}

func (s *UnitService) UpdateUnit(userId, unitId int, input model.UpdateUnitInput) error {
	if err := s.Unit.UnitBelongsToUser(userId, unitId); err != nil {
		return err
	}

	return s.Unit.UpdateUnit(unitId, input)
}

func (s *UnitService) DeleteUnit(userId, unitId int) error {
	if err := s.Unit.UnitBelongsToUser(userId, unitId); err != nil {
		return err
	}

	return s.Unit.DeleteUnit(unitId)
}

func (s *UnitService) ReservedUnits(userId int) ([]model.StorageUnit, error) {
	return s.Unit.ReservedUnits(userId)
}
