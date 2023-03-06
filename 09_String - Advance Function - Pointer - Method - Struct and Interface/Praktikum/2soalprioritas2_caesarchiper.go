package main

import (
	"fmt"
)

func caesar(offset int, input string) string {

	var en string

	for _, b := range []byte(input) {

		n := int(b) - 97 // start at 0
		n += offset
		n %= 26 // length of text ascii lower

		en += string(byte(n + 97))
	}

	return en
}

func main() {

	fmt.Println("Hello World, Hello Golang World!")
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}
