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

// Limits for partnership tiers.
var UnitsPlanLimits = map[int]int{
	//Free tier.
	0: 25,
	//Pro tier.
	1: 100,
	//Enterprise tier.
	2: 32768,
}

type UnitService struct {
	repository.Unit
	repository.Group
	repository.Partnership
}

func NewUnitService(repo *repository.Repository) *UnitService {
	return &UnitService{
		Unit:        repo.Unit,
		Group:       repo.Group,
		Partnership: repo.Partnership,
	}
}

func (s *UnitService) GroupUnits(user m.UserIdentity, groupId int) ([]m.StorageUnit, error) {
	if !s.Group.ManagerOwnsGroup(user.Id, groupId) && user.Role != m.AdminRole {
		return nil, ErrOwnershipViolation
	}

	return s.Unit.GroupUnits(groupId)
}

func (s *UnitService) CreateUnit(user m.UserIdentity, groupId int, unit m.StorageUnit) (int, error) {
	//Check if user has reached limit of units for his partnership tier.
	if err := s.checkUserLimit(user.Id); err != nil {
		return 0, err
	}
	//Check if user is manager of this group.
	if !s.Group.ManagerOwnsGroup(user.Id, groupId) && user.Role != m.AdminRole {
		s.LogHistory(user.Id, groupId, StatusForbidden, UnitCreateAction)
		return 0, ErrOwnershipViolation
	}

	unit.GroupId = groupId
	unit.IsOccupied = false

	return s.Unit.CreateUnit(unit)
}
func (s *UnitService) checkUserLimit(userId int) error {
	part, err := s.Partnership.PartByUserId(userId)
	if err != nil {
		return err
	}
	unitsCount, err := s.Unit.ManagerUnitsCount(userId)
	if err != nil {
		return err
	}
	if unitsCount >= UnitsPlanLimits[part.Tier] {
		return ErrUnitsLimitReached
	}

	return nil
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

	unit, err := s.Unit.UnitById(unitId)
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
	occupied := true
	input := m.UnitInput{
		UserId:     &userId,
		IsOccupied: &occupied,
		BusyUntil:  reservInfo.BusyUntil,
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
