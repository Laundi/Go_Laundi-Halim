package main

import (







	"github.com/labstak/echo/v4/middleware"
)

func main() {
	
	db, err := db.InitDB()
	if err != nil {

		panic (err.Error())
	}

	db.AutoMigrate(&models.User{}, &models.Book{})

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// use echojwt
	// e. Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// SigningKey: utils.JWT_SECRET_KEY
	// }))

	utils.UseStats(e) // Custom Middleware
	e.Use(utils.ServerHeader)

	routers.UserRouter(e, db)
	routers.BookRouther(e. db)

	e. Logger.Fatal(e.Start(":8000"))
}