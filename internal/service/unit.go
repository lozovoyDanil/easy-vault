package service

import (
	"time"

	"main.go/internal/model"
	"main.go/internal/repository"
)

// Action text for logging history
const (
	UnitCreateAction  = "Unit created"
	UnitUpdateAction  = "Unit updated"
	UnitDeleteAction  = "Unit deleted"
	UnitReserveAction = "Unit reserved"
	UnitLockAction    = "Unit locked"
	UnitUnlockAction  = "Unit unlocked"
)

// Status codes for logging history
const (
	StatusFailed = iota
	StatusOK
	StatusForbidden
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
		s.LogHistory(userId, unitId, StatusFailed, UnitUpdateAction)
		return err
	}
	if count == 0 {
		s.LogHistory(userId, unitId, StatusForbidden, UnitUpdateAction)
		return ErrOwnershipViolation
	}

	err = s.Unit.UpdateUnit(unitId, input)
	if err != nil {
		s.LogHistory(userId, unitId, StatusFailed, UnitUpdateAction)
		return err
	}
	s.LogHistory(userId, unitId, StatusOK, UnitUpdateAction)

	return nil
}

func (s *UnitService) DeleteUnit(userId, unitId int) error {
	count, err := s.Unit.UnitBelongsToUser(userId, unitId)
	if err != nil {
		s.LogHistory(userId, unitId, StatusFailed, UnitDeleteAction)
		return err
	}
	if count == 0 {
		s.LogHistory(userId, unitId, StatusForbidden, UnitDeleteAction)
		return ErrOwnershipViolation
	}

	err = s.Unit.DeleteUnit(unitId)
	if err != nil {
		s.LogHistory(userId, unitId, StatusFailed, UnitDeleteAction)
		return err
	}
	s.LogHistory(userId, unitId, StatusOK, UnitDeleteAction)

	return nil
}

func (s *UnitService) ReservedUnits(userId int) ([]model.StorageUnit, error) {
	return s.Unit.ReservedUnits(userId)
}

func (s *UnitService) UnitDetails(userId, unitId int) (model.UnitDetails, error) {
	count, err := s.Unit.UnitBelongsToUser(userId, unitId)
	if err != nil {
		return model.UnitDetails{}, err
	}
	if count == 0 {
		return model.UnitDetails{}, ErrOwnershipViolation
	}

	unit, err := s.UnitById(userId, unitId)
	if err != nil {
		return model.UnitDetails{}, err
	}
	hist, err := s.Unit.UnitHistory(unitId)
	if err != nil {
		return model.UnitDetails{}, err
	}

	return model.UnitDetails{StorageUnit: unit, History: hist}, nil
}

func (s *UnitService) ReserveUnit(userId, unitId int, reservInfo model.UpdateUnitInput) error {
	input := model.UpdateUnitInput{
		UserId:    &userId,
		BusyUntil: reservInfo.BusyUntil,
	}
	err := s.Unit.UpdateUnit(unitId, input)
	if err != nil {
		s.LogHistory(userId, unitId, StatusFailed, UnitReserveAction)
		return err
	}
	s.LogHistory(userId, unitId, StatusOK, UnitReserveAction)

	return nil
}

func (s *UnitService) LogHistory(userId, unitId, status int, action string) error {
	log := model.UnitHistory{
		UnitId:     unitId,
		UserId:     userId,
		Status:     status,
		Action:     action,
		ActionDate: time.Now(),
	}

	return s.Unit.LogHistory(log)
}
