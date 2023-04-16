package main

import "fmt"

func Add(a, b float 64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply (a, b float64) float64 {
	return a * b
}

func Divide (a, b float 64) (float64, error) {
	if b === 0 {
		return 0, fmt.Error("Tidak boleh dibagi nol")
	}
	return a/b, nil
}