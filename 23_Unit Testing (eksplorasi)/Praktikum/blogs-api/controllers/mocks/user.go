package controllers

import (
	"blogs-api/models"
	"blogs-api/services"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	service services.UserService
}

func InitUserController() UserController {
	return UserController{
		service: services.InitUserService(),
	}
}

func (uc *UserController) GetAll(c echo.Context) error {
	users, err := uc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status: "failed",
			Message: "failed to fetch user data",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.User]{
		Status:		"success",
		Message: 	"all users",
		Data:		Users,
	})	
}

func (uc *UserController) GetByID(c echo.Context) error {
	var userID string = c.Param("id")
	
	user, err := uc.service.GetByID(userID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[models.User]{
			Status: "failes",
			Message: "user not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.User]{
		Status:		"success",
		Message:	"user found",
		Data:		User,
	})
}

func (uc *UserController) Create(c. echo.Context) error {
	var userInput models.UserInput

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
				Status: "failed",
				Message: "Invalid request",
		})
	}

	validate := validator.New()
	if err := validate.Struct(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status: "failed",
			Message: err.Error(),
		})
	}

	user, err := uc.service.Service.Create(userInput)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status: "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.User]{
		Status:	"success",
		Message: "user created",
		Data: user,
	})
}

func (uc *UserController) Update(c echo.Context) error {
	var userID string = c.Param("id")

	var userInput models.UserInput

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest) models.Response(string)
		Status: "failed",
		Message: err,Error(),
	}
}