package main

import (
	"fmt"
)

func multipleOf(x int, n int) bool {

	return n % x == 0
}

func main() {

	fmt.Println(multipleOf(7, 21))
	fmt.Println(multipleOf(7, 32))
}
