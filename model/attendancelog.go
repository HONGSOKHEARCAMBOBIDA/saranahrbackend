package models

type AttendanceLog struct {
	ID                int     `json:"id" gorm:"primarykey"`
	EmployeeShiftID   int     `json:"employee_shift_id" gorm:"column:employee_shift_id"`
	CheckDate         string  `json:"check_date" gorm:"column:check_date"`
	CheckIn           string  `json:"check_in" gorm:"column:check_in"`
	CheckOut          *string `json:"check_out" gorm:"column:check_out"`
	Islate            int     `json:"is_late" gorm:"column:is_late"`
	IsLeftEarly       int     `json:"is_left_early" gorm:"column:is_left_early"`
	ISZoonCheckIn     bool    `json:"is_zoon_check_in" gorm:"column:is_zoon_check_in"`
	ISZoonCheckOut    bool    `json:"is_zoon_check_out" gorm:"column:is_zoon_check_out"`
	LatitudeCheckIn   float64 `json:"latitude_check_in" gorm:"column:latitude_check_in"`
	LongitudeCheckIn  float64 `json:"longitude_check_in" gorm:"column:longitude_check_in"`
	LatitudeCheckOut  float64 `json:"latitude_check_out" gorm:"column:latitude_check_out"`
	LongitudeCheckOut float64 `json:"longitude_check_out" gorm:"column:longitude_check_out"`
	Notes             string  `json:"notes"`
	BranchID          int     `json:"branch_id" gorm:"column:branch_id"`
	Status            int     `json:"status" gorm:"column:status"`
	CreateBy          int     `json:"create_by" gorm:"column:create_by"`
}

type AttendanceLogRequestCreate struct {
	EmployeeShiftID int `json:"employee_shift_id" binding:"required"`

	Latitude float64 `json:"latitude" binding:"required"`

	Longitude float64 `json:"longitude" binding:"required"`

	Notes string `json:"notes"`
}
type AttendanceResponse struct {
	ID                int    `json:"id"`
	EmployeeShiftID   int    `json:"employee_shift_id "`
	CheckDate         string `json:"check_date"`
	CheckIn           string `json:"check_in"`
	CheckOut          string `json:"check_out"`
	IsLate            int    `json:"is_late"`
	IsLeftEarly       int    `json:"is_left_early"`
	BranchID          int    `json:"branch_id"`
	BranchName        string `json:"branch_name"`
	Status            int    `json:"status"`
	Nameen            string `json:"name_en" gorm:"column:name_en"`
	NameKh            string `json:"name_kh" gorm:"column:name_kh"`
	RoleID            int    `json:"role_id"`
	RoleName          string `json:"role_name"`
	Type              int    `json:"type" gorm:"column:type"`
	ShiftID           int    `json:"shift_id"`
	ShiftName         string `json:"shift_name"`
	StartTime         string `json:"start_time"`
	EndTime           string `json:"end_time"`
	ISZoonCheckIn     bool   `json:"is_zoon_check_in" gorm:"column:is_zoon_check_in"`
	ISZoonCheckOut    bool   `json:"is_zoon_check_out" gorm:"column:is_zoon_check_out"`
	LatitudeCheckIn   string `json:"latitude_check_in" gorm:"column:latitude_check_in"`
	LongitudeCheckIn  string `json:"longitude_check_in" gorm:"column:longitude_check_in"`
	LatitudeCheckOut  string `json:"latitude_check_out" gorm:"column:latitude_check_out"`
	LongitudeCheckOut string `json:"longitude_check_out" gorm:"column:longitude_check_out"`
	Notes             string `json:"notes"`
}
