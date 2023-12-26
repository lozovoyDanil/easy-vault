package model

type Revenue struct {
	SpacesRevenue []SpaceRevenue
	TotalRevenue  float64
}

type SpaceRevenue struct {
	Name          string
	GroupsRevenue []GroupRevenue
	TotalRevenue  float64
}

type GroupRevenue struct {
	Name         string
	TotalRevenue float64
}
