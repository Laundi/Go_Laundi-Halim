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

var primes = [...]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541}

func isDivAbs(a int, b int) (bool, int) {

	n := a / b // no cast, int type
	k := float64(a) / float64(b)

	return k - float64(n) == 0., n // rest of nums, if gt 0. is no abs
	
}

func getFactorOfNums(x int) []int {

	n := make([]int, 0)
	i := 0 // index of primes
	j := x // rest of values

	var prime int // initial

	// all value can divided by 1
	n = append(n, 1)

	for {

		if j == 1 { break }

		prime = primes[i]

		isAbs, k := isDivAbs(j, prime)

		if isAbs {

			n = append(n, prime)
			j = k
			
		} else { i = i + 1 } // change, next prime value
	}
	return n
}

func main() {

	res, _ := input("Input: ")
	n, _ := strconv.ParseInt(res, 0, 32)

	fmt.Println("Output:")
	fmt.Printf("value: %d\n", n)

	for _, v := range getFactorOfNums(int(n)) {

		fmt.Println(v)
	}
	
	fmt.Println(n)
}
