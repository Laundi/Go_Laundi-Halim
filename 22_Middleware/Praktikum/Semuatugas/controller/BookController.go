package controller

import (
	"go_Laundi-Halim/21_ORM-MVC/Praktikum/Semuatugas/database"
	"go_Laundi-Halim/21_ORM-MVC/Praktikum/Semuatugas/models"
	"net/http"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

var blogs []models.Blog

// get all books
func GetBooksController(c echo.Context) error {
	err := database.DB.Preload("user").Find(&blogs).error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all blogs",
		"Books": books,
	})
}

// get book by id
func GetBooksController(c echo.Context) error {
	blogID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid Blog ID")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog by id",
		"Book":	  books,
}

