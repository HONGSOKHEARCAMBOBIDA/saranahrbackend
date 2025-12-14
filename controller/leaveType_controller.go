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

func CreateLeaveType(c *gin.Context) {
	var input models.LeaveTypeRequest

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

	newleaveType := models.LeaveType{

		Name:         input.Name,
		DeductTypeID: input.DeductTypeID,
		CurrencyID:   input.CurrencyID,
		DeductAmount: input.DeductAmount,
		Description:  input.Description,
		Isactive:     true,
	}

	if err := tx.Create(&newleaveType).Error; err != nil {
		tx.Rollback()
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := tx.Commit().Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	share.ResponeSuccess(c, 200, "Leave Type Create Success")

}

func GetLeaveType(c *gin.Context) {
	var leaveType []models.LeaveTypeResponse

	db := config.DB.Table("leave_types").Select(`
	
		leave_types.id AS id,
		leave_types.name AS name,
		deduct_types.id AS deduct_type_id,
		deduct_types.name AS deduct_type_name,
		currencies.id AS currency_id,
		currencies.code AS currency_code,
		currencies.symbol AS currency_symbol,
		currencies.name AS currency_name,
		leave_types.deduct_amount AS deduct_amount,
		leave_types.description AS description
	`).
		Joins("INNER JOIN deduct_types ON deduct_types.id = leave_types.deduct_type_id").
		Joins("INNER JOIN currencies ON currencies.id = leave_types.currency_id")

	db = db.Order("leave_types.id DESC")

	if err := db.Scan(&leaveType).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	share.RespondDate(c, 200, leaveType)
}

func UpdateLeaveType(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		share.RespondError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var updateleave models.LeaveTypeRequest

	if err := c.ShouldBindJSON(&updateleave); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	result := config.DB.Model(&models.LeaveType{}).Where("id =?", id).Updates(&updateleave)

	if result.RowsAffected == 0 {
		share.RespondError(c, http.StatusBadRequest, "LeaveTypeរកមិនឃេីញឬមិនមានការកែប្រែ")
		return
	}

	share.ResponeSuccess(c, 200, "Leave Type Update Success")

}

func ChangeStatusLeaveType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		share.RespondError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	result := config.DB.Model(&models.LeaveType{}).Where("id =?", id).
		Update("is_active", gorm.Expr("!is_active"))

	if result.Error != nil {
		share.RespondError(c, http.StatusNotFound, result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		share.RespondError(c, http.StatusNotFound, "Leave Type រកមិនឃេីញឬមិនមានការកែប្រែ")
		return
	}

	share.ResponeSuccess(c, http.StatusOK, "Update Status Success")
}
