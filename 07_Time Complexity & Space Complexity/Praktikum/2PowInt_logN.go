package main

import (
	"fmt"
)

func pow(x int, n int) int {

	if n == 0 { return 1 }
	tmp := pow(x, int(n / 2))

	if n % 2 != 0 { return tmp * tmp * x }
	return tmp * tmp
}

func main() {

	fmt.Println(pow(2,3))
	fmt.Println(pow(5,3))
	fmt.Println(pow(10,2))
	fmt.Println(pow(2,5))
	fmt.Println(pow(7,3))
}
