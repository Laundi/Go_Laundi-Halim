package models

import "github.com/jinzu/gorm"

type Book struct {
	gorm.Model
	IDBuku			int		`json:"IdBuku" form:"IdBuku"`
	Judul		string `json: "judul" form: "judul"`
	Penulis 	string `json: "penulis" form:"penulis"`
	Penerbit	string `json: "penerbit" form:"penerbit"`
}
