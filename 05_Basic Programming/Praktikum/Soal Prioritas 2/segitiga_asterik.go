package main

import (
	"fmt"
	"strconv"
)

func input(msg string) (string, error) {

	var result string
	fmt.Print(msg)
	fmt.Scanln(&result)

	return result, nil
}

func main() {

	res, _ := input("Input: ")
	n, _ := strconv.ParseInt(res, 0, 32)

	fmt.Println("Output:")

	var j int
	var m int
	var k int = int(n) // casting, down
	for i := 0; i < k; i++ {

		for j = 0; j < (k - i - 1); j++ {

			fmt.Print(" ")
		}

		m = i + i + 1
		
		for j = 1; j <= m; j++ {

			if j & 1 == 1 {

				fmt.Print("*")
				continue
			}

			fmt.Print(" ")
		}

		fmt.Println()
	}
}
