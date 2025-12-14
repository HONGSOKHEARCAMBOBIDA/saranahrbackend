package controller

import (
	"HRbackend/config"
	"HRbackend/constant/share"
	models "HRbackend/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateDeductType(c *gin.Context) {

	var input models.DeductTypeRequest

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

	newDeducttype := models.DeductType{
		Name:     input.Name,
		IsActive: true,
	}

	if err := tx.Create(&newDeducttype).Error; err != nil {
		tx.Rollback()
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := tx.Commit().Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	share.ResponeSuccess(c, 200, "DeductType Create Success")
}

func GetDeductType(c *gin.Context) {
	var deducttype []models.DeductType

	if err := config.DB.Find(&deducttype).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	share.RespondDate(c, 200, deducttype)
}

func UpdateDeductType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		share.RespondError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var updatededucttype models.DeductTypeRequest

	if err := c.ShouldBindJSON(&updatededucttype); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	result := config.DB.Model(&models.DeductType{}).Where("id =?", id).Updates(updatededucttype)

	if result.RowsAffected == 0 {
		share.RespondError(c, http.StatusBadRequest, "Deducttypeរកមិនឃេីញឬមិនមានការកែប្រែ")
		return
	}

	share.ResponeSuccess(c, 200, "Deducttype Update Success")
}

func ChnageStatusDeductType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		share.RespondError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	result := config.DB.Model(&models.DeductType{}).Where("id =?", id).
		Update("is_active", gorm.Expr("!is_active"))

	if result.Error != nil {
		share.RespondError(c, http.StatusNotFound, result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		share.RespondError(c, http.StatusNotFound, "Deducttypeរកមិនឃេីញឬមិនមានការកែប្រែ")
		return
	}

	share.ResponeSuccess(c, http.StatusOK, "Update Status Success")
}
