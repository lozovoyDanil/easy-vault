package model

type StorageGroup struct {
	Id        int
	Name      string `json:"name"`
	Size      int    `json:"size"`
	NumOfFree int    `json:"numOfFree" db:"numOfFree"`
}
