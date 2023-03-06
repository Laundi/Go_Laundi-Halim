package main

import (

	"fmt"
)

func factorial(num int) int {

	if num <= 1 { return 1 }
	
	return num * factorial(num - 1)
}

func pascal(num int) []int {

	temp := make([]int, 0)

	for i := 0; i < num + 1; i++ {

		a := factorial(num)
		b := factorial(num - i)
		c := factorial(i)
		k := b * c

		temp = append(temp, int(a / k))
	}

	return temp
}

func pascals(num int) [][]int {

	temp := make([][]int, 0)

	for i := 0; i < num; i++ {

		temp = append(temp, pascal(i))
	}

	return temp
}

func main() {

	fmt.Println(pascals(5))
}
