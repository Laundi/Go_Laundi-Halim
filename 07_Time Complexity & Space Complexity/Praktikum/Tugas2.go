package main

import (
	"fmt"
)

func pow (x, n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	halfpow := pow(x, n/2)
	result := halfpow * halfpow

	if n%2 != 0
			result *= x
	}

	return result


func main() {
	fmt.Println(pow(2,3))
	fmt.Println(pow(5,3))
	fmt.Println(pow(10,2))
	fmt.Println(pow(2,5))
	fmt.Println(pow(7,3))
}