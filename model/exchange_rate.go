package models

type ExchangeRate struct {
	ID       int     `json:"id" gorm:"primarykey"`
	PairID   int     `json:"pair_id" gorm:"column:pair_id"`
	Rate     float64 `json:"rate" gorm:"column:rate"`
	Isactive int     `json:"is_active" gorm:"column:is_active"`
	IsEdit   int     `json:"is_edit" gorm:"column:is_edit"`
	CreateBY int     `json:"create_by" gorm:"column:create_by"`
	UpdateBy int     `json:"update_by" gorm:"column:update_by"`
}
type ExchangeRateRequest struct {
	PairID   int     `json:"pair_id" gorm:"column:pair_id"`
	Rate     float64 `json:"rate" gorm:"column:rate"`
	Isactive int     `json:"is_active" gorm:"column:is_active"`
	IsEdit   int     `json:"is_edit" gorm:"column:is_edit"`
	CreateBY int     `json:"create_by" gorm:"column:create_by"`
	UpdateBy int     `json:"update_by" gorm:"column:update_by"`
}
type ExchangeRateResponse struct {
	ID                     int    `json:"id"`
	PairID                 int    `json:"pair_id"`
	BaseCurrencyID         int    `json:"base_currency_id"`
	BaseCurrencyCode       string `json:"base_currency_code"`
	BaseCurrencySymbol     string `json:"base_currency_symbol"`
	BaseCurrencyName       string `json:"base_currency_name"`
	BaseCurrencyIsactive   bool   `json:"base_currency_is_active" gorm:"column:base_currency_is_active"`
	TargetCurrencyID       int    `json:"target_currency_id"`
	TargetCurrencyCode     string `json:"target_currency_code"`
	TargetCurrencySymbol   string `json:"target_currency_symbol"`
	TargetCurrencyName     string `json:"target_currency_name"`
	TargetCurrencyIsactive bool   `json:"target_currency_is_active" gorm:"column:target_currency_is_active"`
	Rate                   string `json:"rate"`
	Isactive               int    `json:"is_active" gorm:"column:is_active"`
	IsEdit                 int    `json:"is_edit"`
	CreateBy               int    `json:"create_by"`
	CreateByName           string `json:"create_by_name"`
	UpdateBy               int    `json:"update_by"`
	UpdateByName           string `json:"update_by_name"`
}
