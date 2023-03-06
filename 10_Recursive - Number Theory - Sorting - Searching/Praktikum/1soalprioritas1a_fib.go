package main

import (
	"fmt"
)

func fib(num int) int {

	if num <= 1 { return num }
	return fib(num - 1) + fib(num - 2)
}

func main() {

	fmt.Println(fib(0))
	fmt.Println(fib(2))
	fmt.Println(fib(9))
	fmt.Println(fib(10))
	fmt.Println(fib(12))
}
