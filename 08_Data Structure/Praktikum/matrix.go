package main

import (
	"fmt"
)

func Abs(a int, b int) int {

	if a <= b { return b - a }
	return a - b
}

func AbsDiffArrMatrix(n int, data [][]int) int {
	
	var a, b int

	for i, v := range data {

		// a
		a += v[i]
		
		// b
		b += v[n - i - 1]
	}

	return Abs(a, b)
}

func main() {

	data := [][]int{[]int{1,2,3},[]int{4,5,6},[]int{9,8,9}}

	fmt.Println(AbsDiffArrMatrix(3, data))
}
