package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {

	x := len(a)
	y := len(b)

	if x < y {

		if 0 <= strings.Compare(b, a) {

			return a
		}
		
	} else {

		if 0 <= strings.Compare(a, b) {

			return b
		}
	}

	return ""
}

func main() {

	fmt.Println("Hello World, Hello Golang World!")
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}
