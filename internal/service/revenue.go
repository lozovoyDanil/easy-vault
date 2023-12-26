package service

import (
	m "main.go/internal/model"
	"main.go/internal/repository"
)

type RevenueService struct {
	repository.Partnership
	repository.Space
	repository.Group
	repository.Unit
}

func NewRevenueService(repo *repository.Repository) *RevenueService {
	return &RevenueService{
		Partnership: repo.Partnership,
		Space:       repo.Space,
		Group:       repo.Group,
		Unit:        repo.Unit,
	}
}

func (s *RevenueService) ManagerRevenue(userId int) (m.Revenue, error) {
	var revenue m.Revenue
	part, err := s.Partnership.PartByUserId(userId)
	if err != nil {
		return revenue, err
	}
	if part.Tier != EnterpriseTier {
		return revenue, ErrTierViolation
	}

	spaces, err := s.Space.UserSpaces(userId)
	if err != nil {
		return revenue, err
	}
	groups, err := s.userGroups(spaces)
	if err != nil {
		return revenue, err
	}
	units, err := s.userUnits(groups)
	if err != nil {
		return revenue, err
	}
	revenue = s.calculateRevenue(spaces, groups, units)

	return revenue, nil
}

func (s *RevenueService) calculateRevenue(
	spaces []m.Space,
	groups []m.StorageGroup,
	units []m.StorageUnit,
) m.Revenue {
	groupMap := make(map[int]*m.StorageGroup)
	for i := range groups {
		groupMap[groups[i].Id] = &groups[i]
	}

	spaceMap := make(map[int]*m.SpaceRevenue)
	for _, space := range spaces {
		spaceMap[space.Id] = &m.SpaceRevenue{Name: space.Name}
	}

	for _, unit := range units {
		if !unit.IsOccupied {
			continue
		}

		group := groupMap[unit.GroupId]
		spaceRevenue := spaceMap[group.SpaceId]
		found := false
		for i, groupRevenue := range spaceRevenue.GroupsRevenue {
			if groupRevenue.Name == group.Name {
				spaceRevenue.GroupsRevenue[i].TotalRevenue += float64(group.Price)
				found = true
				break
			}
		}
		if !found {
			spaceRevenue.GroupsRevenue = append(spaceRevenue.GroupsRevenue, m.GroupRevenue{
				Name:         group.Name,
				TotalRevenue: float64(group.Price),
			})
		}
		spaceRevenue.TotalRevenue += float64(group.Price)
	}

	var totalRevenue float64
	var revenueData m.Revenue
	for _, spaceRevenue := range spaceMap {
		revenueData.SpacesRevenue = append(revenueData.SpacesRevenue, *spaceRevenue)
		totalRevenue += spaceRevenue.TotalRevenue
	}
	revenueData.TotalRevenue = totalRevenue

	return revenueData
}

func (s *RevenueService) userGroups(spaces []m.Space) ([]m.StorageGroup, error) {
	var groups []m.StorageGroup
	for _, space := range spaces {
		spaceGroups, err := s.Group.SpaceGroups(space.Id)
		if err != nil {
			return nil, err
		}
		groups = append(groups, spaceGroups...)
	}

	return groups, nil
}

func (s *RevenueService) userUnits(groups []m.StorageGroup) ([]m.StorageUnit, error) {
	var units []m.StorageUnit
	for _, group := range groups {
		groupUnits, err := s.Unit.GroupUnits(group.Id)
		if err != nil {
			return nil, err
		}
		units = append(units, groupUnits...)
	}

	return units, nil
}
