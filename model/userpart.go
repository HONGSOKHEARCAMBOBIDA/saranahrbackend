package models

type UserPart struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	PartID int `json:"part_id"`
}

type UserPartRequest struct {
	UserID int `json:"user_id"`
	PartID int `json:"part_id"`
}
