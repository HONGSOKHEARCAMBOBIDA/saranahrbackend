package controller

import (
	"HRbackend/config"
	"HRbackend/constant/share"
	"HRbackend/helper"
	models "HRbackend/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateExchangeRate(c *gin.Context) {
	var input models.ExchangeRateRequest
	userid, ok := helper.GetUserID(c)
	if !ok {
		share.RespondError(c, http.StatusUnauthorized, "Please Login")
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	tx := config.DB.Begin()
	if tx.Error != nil {
		share.RespondError(c, http.StatusInternalServerError, tx.Error.Error())
		return
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}

	}()
	newExchangeRate := models.ExchangeRate{
		PairID:   input.PairID,
		Rate:     input.Rate,
		Isactive: 1,
		IsEdit:   0,
		CreateBY: userid,
		UpdateBy: userid,
	}
	if err := tx.Create(&newExchangeRate).Error; err != nil {
		tx.Rollback()
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	tx.Commit()
	share.ResponeSuccess(c, http.StatusOK, "create success")
}

func GetExchangeRate(c *gin.Context) {
	var ExchangeRate []models.ExchangeRateResponse
	db := config.DB.Table("exchange_rates").Select(`
	exchange_rates.id AS id,
	b.id AS base_currency_id,
	b.code AS base_currency_code,
	b.symbol AS base_currency_symbol,
	b.name AS base_currency_name,
	b.is_active AS base_currency_is_active,
	t.id AS target_currency_id,
	t.code AS target_currency_code,
	t.symbol AS target_currency_symbol,
	t.name AS target_currency_name,
	t.is_active AS target_currency_is_active,
	exchange_rates.rate AS rate,
	exchange_rates.is_active AS is_active,
	exchange_rates.is_edit AS is_edit,
	exchange_rates.pair_id AS pair_id,
	ec.id AS create_by,
	ec.name_kh AS create_by_name,
	eu.id AS update_by,
	eu.name_kh AS update_by_name
	
	`).
		Joins("INNER JOIN currency_pairs c ON c.id = exchange_rates.pair_id").
		Joins("INNER JOIN currencies b ON b.id = c.base_currency_id").
		Joins("INNER JOIN currencies t ON t.id = c.target_currency_id").
		Joins("INNER JOIN users u ON u.id = exchange_rates.create_by").
		Joins("INNER JOIN employees ec ON ec.id = u.employee_id").
		Joins("INNER JOIN users ON users.id = exchange_rates.update_by").
		Joins("INNER JOIN employees eu ON eu.id = users.employee_id").
		Order("exchange_rates.id desc")
	if err := db.Find(&ExchangeRate).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	share.RespondDate(c, http.StatusOK, ExchangeRate)
}

func UpdateExchageRate(c *gin.Context) {
	userid, ok := helper.GetUserID(c)
	if !ok {
		share.RespondError(c, http.StatusUnauthorized, "Please Login")
		return
	}

	id := c.Param("id")

	var updateexchangerate models.ExchangeRateRequest
	if err := c.ShouldBindJSON(&updateexchangerate); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Add fixed fields manually because Updates(updateexchangerate)
	// will NOT include fields that are zero-value
	data := map[string]interface{}{
		"pair_id":   updateexchangerate.PairID,
		"rate":      updateexchangerate.Rate,
		"is_active": 1,
		"is_edit":   1,
		"update_by": userid,
	}

	result := config.DB.Model(&models.ExchangeRate{}).
		Where("id = ?", id).
		Updates(data)

	if result.Error != nil {
		share.RespondError(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		share.RespondError(c, http.StatusNotFound, "កែមិនបាន")
		return
	}

	share.ResponeSuccess(c, http.StatusOK, "update success")
}

func ChangeStatusExchangeRate(c *gin.Context) {
	id := c.Param("id")
	result := config.DB.Model(&models.ExchangeRate{}).Where("id =?", id).Update("is_active", gorm.Expr("1 - is_active"))
	if result.Error != nil {
		share.RespondError(c, http.StatusInternalServerError, result.Error.Error())
		return
	}
	share.ResponeSuccess(c, http.StatusOK, "បានប្ដូស្ថានភាព")
}
