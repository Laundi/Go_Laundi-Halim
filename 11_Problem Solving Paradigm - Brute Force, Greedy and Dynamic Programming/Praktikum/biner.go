package main

import (

	"fmt"
)

func pow(a int, b int) int {

	if b < 1 { return 1 }
	return a * pow(a, b - 1)
}

func bin(num int) string {

	if num == 0 { return "0" }

	var temp string
	var k int
	
	i := 0

	for {

		k = pow(2, i)

		if num < k { break }
		
		i += 1
		
		if num & k != 0 {

			temp = "1" + temp
			continue
		}

		temp = "0" + temp
	}

	return temp 
}

func binaries(num int) []string {

	temp := make([]string, 0)

	for i := 0; i < num + 1; i++ {

		temp = append(temp, bin(i))
	}

	return temp
}

func main() {

	fmt.Println(binaries(2))
	fmt.Println(binaries(3))
}
