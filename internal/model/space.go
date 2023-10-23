package model

type Space struct {
	Id          int    `json:"-" db:"id"`
	Name        string `json:"name"`
	Addr        string `json:"addr"`
	NumOfGroups int    `json:"numOfGroups" db:"numOfGroups"`
	Size        int    `json:"size"`
	NumOfFree   int    `json:"numOfFree" db:"numOfFree"`
}

type UpdateSpaceInput struct {
	Name *string `json:"name"`
	Addr *string `json:"addr"`
}
