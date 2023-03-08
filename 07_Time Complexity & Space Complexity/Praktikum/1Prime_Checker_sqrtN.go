package main

import (
	"fmt"
)

func primeChecker(num int) bool {

	if num < 2 { return false } // 0, 1
	if num == 2 { return true } // 2
	if num & 1 == 0 { return false } // evenly
	
	for i := 3; i * i <= num; i++ {

		// if num & 1 == 0 { continue } // after 2, most prime numbers in odd numbers
		if num % i == 0 { return false } // can`t be divided by the value below
		i += 1; // odd
	}
	
	// if num % 5 == 0 { return false } // modulo 5
// 
	// i := int(num / 2)
// 
	// for i > 2 {
// 
		// if num % i == 0 { return false }
		// i = int(i / 2)
// 
	// }
	
	return true 
}

func main() {

	fmt.Println(primeChecker(1000000007)) // prime
	fmt.Println(primeChecker(13)) // prime
	fmt.Println(primeChecker(17)) // prime
	fmt.Println(primeChecker(20)) // not prime
	fmt.Println(primeChecker(35)) // not prime
	fmt.Println(primeChecker(11 * 13)) // not prime
	fmt.Println(primeChecker(1500450271)) // prime
}
