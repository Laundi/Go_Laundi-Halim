package main

import (
	"fmt"
)

type Car struct {

	Type string `json:"type"`
	FullIn bool `json:"full"`
}

func NewCar() *Car {

	car := Car{"Avanza", true}
	return &car
}

func (this *Car)GetDistanceInMill(benzin float64) float64 {

	return benzin / 1.5
}

func main() {

	fmt.Println("Hello World, Hello Golang World!")

	car := NewCar()
	fmt.Println(car.GetDistanceInMill(22.5))
}
