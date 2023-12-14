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
	return &UnitService{Unit: repo.Unit, Group: repo.Group}
}

func (s *UnitService) GroupUnits(userId, groupId int) ([]model.StorageUnit, error) {
	count, err := s.Group.GroupBelongsToUser(userId, groupId)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, ErrOwnershipViolation
	}

	return s.Unit.GroupUnits(groupId)
}

func (s *UnitService) UnitById(userId, unitId int) (model.StorageUnit, error) {
	count, err := s.Unit.UnitBelongsToUser(userId, unitId)
	if err != nil {
		return model.StorageUnit{}, err
	}
	if count == 0 {
		return model.StorageUnit{}, ErrOwnershipViolation
	}

	return s.Unit.UnitById(unitId)
}

func (s *UnitService) CreateUnit(userId, groupId int, unit model.StorageUnit) (int, error) {
	count, err := s.Group.GroupBelongsToUser(userId, groupId)
	if err != nil {
		return 0, err
	}
	if count == 0 {
		return 0, ErrOwnershipViolation
	}

	unit.GroupId = groupId

	return s.Unit.CreateUnit(unit)
}

func (s *UnitService) UpdateUnit(userId, unitId int, input model.UpdateUnitInput) error {
	count, err := s.Unit.UnitBelongsToUser(userId, unitId)
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrOwnershipViolation
	}

	return s.Unit.UpdateUnit(unitId, input)
}

func (s *UnitService) DeleteUnit(userId, unitId int) error {
	count, err := s.Unit.UnitBelongsToUser(userId, unitId)
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrOwnershipViolation
	}

	return s.Unit.DeleteUnit(unitId)
}

func (s *UnitService) ReservedUnits(userId int) ([]model.StorageUnit, error) {
	return s.Unit.ReservedUnits(userId)
}

func (s *UnitService) UnitDetails(userId, unitId int) (model.StorageUnit, error) {
	count, err := s.Unit.UnitBelongsToUser(userId, unitId)
	if err != nil {
		return model.StorageUnit{}, err
	}
	if count == 0 {
		return model.StorageUnit{}, ErrOwnershipViolation
	}

	return s.Unit.UnitDetails(unitId)
}

func (s *UnitService) ReserveUnit(userId, unitId int, reservInfo model.UpdateUnitInput) error {
	input := model.UpdateUnitInput{
		UserId:    &userId,
		BusyUntil: reservInfo.BusyUntil,
	}
	err := s.Unit.UpdateUnit(unitId, input)

	return err
}
