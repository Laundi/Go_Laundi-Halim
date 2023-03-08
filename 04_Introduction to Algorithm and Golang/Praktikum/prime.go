package main

import (
	"fmt"
)

func prime(n int) []int {
	x := make([]int, 0)
	b := false
	j := 2
	for {
		if len(x) == n { break }
		for _, v := range(x) {
			if j % v == 0 {
				b = true
				break
			}
		}
		if b { b = false } else { x = append(x, j) }
		j++
	}
	return x
}

func main() {
	for _, v := range(prime(100)) {
		fmt.Printf("%d, ", v)
	}
	fmt.Println()
}
