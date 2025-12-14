package models

type DeductType struct {
	ID       int    `json:"id" gorm:"primarykey"`
	Name     string `json:"name" gorm:"column:name"`
	IsActive bool   `json:"is_active" gorm:"column:is_active"`
}
type DeductTypeRequest struct {
	Name string `json:"name" gorm:"column:name"`
}
