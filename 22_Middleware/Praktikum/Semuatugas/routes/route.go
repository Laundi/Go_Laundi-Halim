package routes

import (
	"go_Laundi-Halim/21_ORM-MVC/Praktikum/Semuatugas/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e : echo.New()

	// routing with query parameter
	e.GET("/users", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.GET("/users/:id", controller.GetUserController)
	e.DELETE("/users/:id", controller.UpdateUserController)
	e.PUT("/users/:id", controller.UpdateUserController)
	// routing with query parameter
	e.GET("/books", controller.GetBookController)
	e.POST("/books", controller.CreateBookController)
	e.GET("/books/:id", controller.GetBookController)
	e.DELETE("/books/:id", controller.Delete.BookController)
	e.PUT("/books/:id", controller.UpdateBookController)
	// routing with query parameter
	e.GET("/blogs", controller.GetBlogsController)
	e.POST("/blogs", controller.CreateBlogController)
	e.GET("/blogs/:id", controller.GetBlogController)
	e.DELETE("/blogs/:id", controller.DeleteBlogController)
	e.PUT("/blogs/:id", controller.UpdateBlogController)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))	
}