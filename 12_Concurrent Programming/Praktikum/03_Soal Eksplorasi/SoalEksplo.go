package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Title		string 	`json:"title"`
	Category	string 	`json:"category"`
	Price		float64 `json:"price"`
	Rating		struct { Rate float64 `json:"rate"`} `json:"rating"`
}

func getProducts(url string, ch chan<- []Product) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error get products :", err)
		return
	}
	defer res.Body.Close()

	var products []Product
	err = json.NewDecoder(res.Body). Decode(&products)
	if err != nil {
		fmt.Println("Error decoding products:", err)
		return
	}

	if len(products) > 5 {
		products = products[:5]
	}

	ch <- products
}

func main() {
	url := "https://fakestoreapi.com/products"
	ch := make(chan []Product)

	go getProducts(url, ch)

	products := <-ch
	fmt.Println("Data Product")
	for _, p := range products {
		fmt.Println("===")
		fmt.Println("Title :", p.Title)
		fmt.Println("Price : $", p.Price)
		fmt.Println("Category :", p.Category)
		fmt.Println("Rating : *", p. Rating.Rate)
		fmt.Println("===")
	}
}
