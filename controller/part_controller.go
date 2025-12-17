package controller

import (
	"HRbackend/config"
	"HRbackend/constant/share"
	models "HRbackend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePart(c *gin.Context) {

	var input models.PartResquest
	if err := c.ShouldBindJSON(&input); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	newpart := models.Part{
		Name: input.Name,
	}
	if err := config.DB.Create(&newpart).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	share.ResponeSuccess(c, 200, "Part Create Success")
}
func Getpart(c *gin.Context) {
	var part []models.Part
	if err := config.DB.Find(&part).Error; err != nil {
		share.RespondError(c, http.StatusNotFound, err.Error())
		return
	}
	share.RespondDate(c, 200, part)
}
