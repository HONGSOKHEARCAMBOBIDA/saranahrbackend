package models

type Part struct {
	ID   int    `json:"id" gorm:"column:id"`
	Name string `json:"name"`
}

type PartResquest struct {
	Name string `json:"name"`
}

type UserPartResponse struct {
	ID       int    `json:"id" gorm:"column:id"`
	PartID   int    `json:"part_id"`
	PartName string `json:"part_name"`
}
