package service

import (
	"time"

	m "main.go/internal/model"
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

func (s *UnitService) GroupUnits(user m.UserIdentity, groupId int) ([]m.StorageUnit, error) {
	if !s.Group.ManagerOwnsGroup(user.Id, groupId) && user.Role != m.AdminRole {
		return nil, ErrOwnershipViolation
	}

	return s.Unit.GroupUnits(groupId)
}

// TODO: delete this function, use unitdetails instead
func (s *UnitService) UnitById(userId, unitId int) (m.StorageUnit, error) {
	ownerId, err := s.Unit.UnitOwnerId(unitId)
	if err != nil {
		return m.StorageUnit{}, err
	}
	if ownerId != userId {
		return m.StorageUnit{}, ErrOwnershipViolation
	}

	return s.Unit.UnitById(unitId)
}

func (s *UnitService) CreateUnit(user m.UserIdentity, groupId int, unit m.StorageUnit) (int, error) {
	if !s.Group.ManagerOwnsGroup(user.Id, groupId) && user.Role != m.AdminRole {
		s.LogHistory(user.Id, groupId, StatusForbidden, UnitCreateAction)
		return 0, ErrOwnershipViolation
	}

	unit.GroupId = groupId

	return s.Unit.CreateUnit(unit)
}

func (s *UnitService) UpdateUnit(user m.UserIdentity, unitId int, input m.UnitInput) error {
	if !s.Unit.ManagerOwnsUnit(user.Id, unitId) && user.Role != m.AdminRole {
		s.LogHistory(user.Id, unitId, StatusForbidden, UnitUpdateAction)
		return ErrOwnershipViolation
	}

	err := s.Unit.UpdateUnit(unitId, input)
	if err != nil {
		s.LogHistory(user.Id, unitId, StatusFailed, UnitUpdateAction)
		return err
	}
	s.LogHistory(user.Id, unitId, StatusOK, UnitUpdateAction)

	return nil
}

func (s *UnitService) DeleteUnit(user m.UserIdentity, unitId int) error {
	if !s.Unit.ManagerOwnsUnit(user.Id, unitId) && user.Role != m.AdminRole {
		s.LogHistory(user.Id, unitId, StatusForbidden, UnitDeleteAction)
		return ErrOwnershipViolation
	}

	err := s.Unit.DeleteUnit(unitId)
	if err != nil {
		s.LogHistory(user.Id, unitId, StatusFailed, UnitDeleteAction)
		return err
	}
	s.LogHistory(user.Id, unitId, StatusOK, UnitDeleteAction)

	return nil
}

func (s *UnitService) ReservedUnits(userId int) ([]m.StorageUnit, error) {
	return s.Unit.ReservedUnits(userId)
}

func (s *UnitService) UnitDetails(user m.UserIdentity, unitId int) (m.UnitDetails, error) {
	ownerId, err := s.Unit.UnitOwnerId(unitId)
	if err != nil {
		return m.UnitDetails{}, err
	}
	if ownerId != user.Id && user.Role == m.CustomerRole {
		return m.UnitDetails{}, ErrOwnershipViolation
	}

	unit, err := s.UnitById(user.Id, unitId)
	if err != nil {
		return m.UnitDetails{}, err
	}
	hist, err := s.Unit.UnitHistory(unitId)
	if err != nil {
		return m.UnitDetails{}, err
	}

	return m.UnitDetails{StorageUnit: unit, History: hist}, nil
}

func (s *UnitService) ReserveUnit(userId, unitId int, reservInfo m.UnitInput) error {
	input := m.UnitInput{
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
	log := m.UnitHistory{
		UnitId:     unitId,
		UserId:     userId,
		Status:     status,
		Action:     action,
		ActionDate: time.Now(),
	}

	return s.Unit.LogHistory(log)
}
