package controllers

import (
	"blogs-api/models"
	"blogs-api/services"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type BookController struct {
	service services.BookService
}

func InitBookController() BookController {
	return BookController{
		service: services.InitBookService(),
	}
}