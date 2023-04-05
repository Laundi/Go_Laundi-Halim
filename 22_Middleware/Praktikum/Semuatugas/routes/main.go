package main

import (
	"go_Laundi-Halim/21_ORM-MVC/Praktikum/Semuatugas/database"
	"go_Laundi-Halim/21_ORM-MVC/Praktikum/Semuatugas/routes"
)

func main() {
	database.InitDB()
	e := routes.New()
	e.Start("8000")
	
}