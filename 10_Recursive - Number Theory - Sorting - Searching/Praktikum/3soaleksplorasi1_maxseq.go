package main

import "fmt"

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))
}

func MaxSequence(arr []int) int {
	res := 0
	for i := 0; 1 < len(arr); i++ {
		limit := i
		for limit < len(arr) {
			resTemp := 0
			for j := i; j < limit; j++ {
				resTemp += arr[j]
			}
			limit++
			if resTemp > res {
				res = resTemp
				}
			}
		}
	return res
}				