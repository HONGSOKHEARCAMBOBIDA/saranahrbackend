package models

type LeaveType struct {
	ID           int     `json:"id" gorm:"primarykey"`
	Name         string  `json:"name"`
	DeductTypeID int     `json:"deduct_type_id"`
	CurrencyID   int     `json:"currency_id"`
	DeductAmount float64 `json:"deduct_amount"`
	Description  string  `json:"description"`
	Isactive     bool    `json:"is_active" gorm:"column:is_active"`
}

type LeaveTypeRequest struct {
	Name         string  `json:"name"`
	DeductTypeID int     `json:"deduct_type_id"`
	CurrencyID   int     `json:"currency_id"`
	DeductAmount float64 `json:"deduct_amount"`
	Description  string  `json:"description"`
}

type LeaveTypeResponse struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	DeductTypeID   int     `json:"deduct_type_id"`
	DeductTypeName string  `json:"deduct_type_name"`
	CurrencyID     int     `json:"currency_id"`
	CurrencyCode   string  `json:"currency_code"`
	CurrencySymbol string  `json:"currency_symbol"`
	CurrencyName   string  `json:"currency_name"`
	DeductAmount   float64 `json:"deduct_amount"`
	Description    string  `json:"description"`
}
