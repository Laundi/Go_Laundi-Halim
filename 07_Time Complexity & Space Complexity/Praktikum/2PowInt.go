package main

import (
	"fmt"
)

func powInt(a int, b int) int {

	result := 1
	for i := 0; i < b; i++ {

		result *= a
	}

	return result
}

func main() {

	fmt.Println(powInt(2,3))
	fmt.Println(powInt(5,3))
	fmt.Println(powInt(10,2))
	fmt.Println(powInt(2,5))
	fmt.Println(powInt(7,3))
}
