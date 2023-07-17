package controllers

import (
	"net/http"
	"pasar/config"
	"pasar/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetPasar(c echo.Context) error {
	var pasar []models.Pasar
	err := config.DB.Find(&pasar).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, pasar)

}

func PostPasar(c echo.Context) error {
	pasar := models.Pasar{}
	c.Bind(&pasar)

	result := config.DB.Create(&pasar)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal memasukkan data", Status: false, Data: nil,
		})
	}
	return c.JSON(http.StatusCreated, models.BaseResponse{
		Message: "Success memasukkan data", Status: true, Data: pasar,
	})
}

func DetailPasar(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid pasar ID")
	}

	var pasar models.Pasar
	err = config.DB.First(&pasar, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, "Pasar not found")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, pasar)
}

func UpdatePasar(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid pasar ID")
	}

	var pasar models.Pasar
	err = config.DB.First(&pasar, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, "Pasar not found")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := c.Bind(&pasar); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	err = config.DB.Save(&pasar).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, pasar)
}

func DeletePasar(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid pasar ID")
	}

	err = config.DB.Delete(&models.Pasar{}, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
