package models

type User struct {
	ID uint `json:"id"`

	BranchID int `json:"branch_id"`

	UserName string `json:"username" gorm:"column:username"`

	Email string `json:"email"`

	Password string `json:"password"`

	Contact string `json:"contact"`

	RoleID int `json:"role_id"`

	EmployeeID int `json:"employee_id"`
}

type LoginReq struct {
	Contact string `json:"contact" bingding:"required"`

	Password string `json:"password" bingding:"required"`
}

type UserReqInsert struct {
	BranchID int `form:"branch_id" binding:"required"`

	NameEn string `form:"name_en" binding:"required"`

	NameKh string `form:"name_kh" binding:"required"`

	UserName string `form:"username"`

	Email string `form:"email" binding:"required,email"`

	Password string `form:"password" binding:"required"`

	Gender int `form:"gender" binding:"required"`

	Contact string `form:"contact" binding:"required"`

	NationalIDNumber string `form:"national_id_number" binding:"required"`

	RoleID int `form:"role_id" binding:"required"`

	EmployeeID int `form:"employee_id"`

	HireDate string `form:"hire_date" binding:"required"`

	PromoteDate string `form:"promote_date" binding:"required"`

	Type int `form:"type" binding:"required"`

	ShiftID int `form:"shift_id" binding:"required"`

	BaseSalary float64 `form:"base_salary" binding:"required"`

	WorkedDay int `form:"worked_day" binding:"required"`

	EffectTiveDate string `form:"effective_date"`

	DateOfBirth string `form:"date_of_birth" binding:"required"`

	VillageIDOfBirht int `form:"village_id_of_birth" binding:"required"`

	MaterialStatus int `form:"marital_status" binding:"required"`

	VillageIDCurrentAddress int `form:"village_id_current_address" binding:"required"`

	FamilyPhone string `form:"family_phone" binding:"required"`

	EducationLevel string `form:"education_level" binding:"required"`

	ExperienceYear int `form:"experience_years" binding:"required"`

	PreviousComapy string `form:"previous_company" binding:"required"`

	BankName string `form:"bank_name" binding:"required"`

	BankAccountNumber string `form:"bank_account_number" binding:"required"`

	PositionLevel int `form:"position_level" binding:"required"`

	Note string `form:"note"`

	CurrencyID int `form:"currency_id" binding:"required"`

	PartIDs []int `form:"part_ids" binding:"required"`
}

type UserReqUpdate struct {
	BranchID int `json:"branch_id"`

	UserName string `json:"username"`

	Email string `json:"email"`

	Contact string `json:"contact"`

	RoleID int `json:"role_id"`

	PartIDs []int `json:"part_ids"`
}

type UserResponse struct {
	Id int `json:"id" gorm:"primarykey"`

	BranchID int `json:"branch_id"`

	BranchName string `json:"branch_name"`

	NameEn string `json:"name_en"`

	NameKh string `json:"name_kh"`

	UserName string `json:"username" gorm:"column:username"`

	Email string `json:"email"`

	Gender int `json:"gender"`

	Contact string `json:"contact"`

	NationalIDNumber string `json:"national_id_number"`

	RoleID int `json:"role_id"`

	RoleName string `json:"role_name"`

	IsActive bool `json:"is_active"`

	UserPartResponse []UserPartResponse `json:"parts" gorm:"-"`
}

type UserResponseV2 struct {
	ID      int     `json:"id"`
	NameKh  string  `json:"name_kh"`
	Branch  Branch  `json:"branch"`
	Role    Role    `json:"role"`
	Village Village `json:"village"`
}
