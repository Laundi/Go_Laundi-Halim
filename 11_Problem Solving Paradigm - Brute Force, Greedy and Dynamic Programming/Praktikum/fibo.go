package main

import (

	"fmt"
)

// func zeros(num int) []int {
// 
	// temp := make([]int, 0)
// 
	// for i := 0; i < num; i++ {
// 
		// temp = append(temp, 0)
	// }
// 
	// return temp
// }

func fibo(num int) int {

	if num == 0 { return 0 }

	// temp := zeros(num + 1)
	temp := make([]int, num + 1)

	temp[0] = 0
	temp[1] = 1

	for i := 2; i < num + 1; i++ {

		temp[i] = temp[i - 1] + temp[i - 2]
	}

	return temp[num]
}

func main() {

	fmt.Println(fibo(0))
	fmt.Println(fibo(1))
	fmt.Println(fibo(2))
	fmt.Println(fibo(3))
	fmt.Println(fibo(4))
	fmt.Println(fibo(5))
	fmt.Println(fibo(6))
	fmt.Println(fibo(7))
	fmt.Println(fibo(8))
	fmt.Println(fibo(9))
	fmt.Println(fibo(10))
}
