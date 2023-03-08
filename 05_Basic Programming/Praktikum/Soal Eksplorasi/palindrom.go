package main

import (
	"fmt"
	"bufio"
	"errors"
	"os"
)

func input(msg string) (string, error) {

	fmt.Print(msg)

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {

		err := scanner.Err()

		if err != nil {
			
			return "", errors.New("Something error!")
		}
	}

	result := scanner.Text()

	return result, nil
}

var lowerCaseLetters = [...]byte{97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122}
var upperCaseLetters = [...]byte{65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90}

func isAlpha(v byte) bool {

	

	// check uppercase letters
	for _, c := range upperCaseLetters {

		if c == v { return true }
	}

	// check lowercase letters
	for _, c := range lowerCaseLetters {

		if c == v { return true }
	}

	// maybe not
	return false
}

func lower(v string) string {

	result := ""
	buff := []byte(v)
	isUpper := false

	for _, c := range buff {

		for i, x := range upperCaseLetters {

			if c == x { 

				result = result + string(lowerCaseLetters[i])
				isUpper = true
				break
			}
		}

		if !isUpper {

			result = result + string(c)
			
		} else {

			isUpper = false // reset
		}
	}

	return result
}

func similarAlpha(a byte, b byte) bool {

	// a - z ; 97 - 122
	// A - Z ; 65 - 90
	// range 32

	if (isAlpha(a) && isAlpha(b)) {

		if lower(string(a)) == lower(string(b)) {

			return true
		} 
	}

	return false
}

func palindrome(context string) bool {

	buff := []byte(context)
	n := len(buff)
	k := n / 2
	j := 0

	var a, b byte

	for i := 0; i < k; i++ {

		j = n - i - 1

		a, b = buff[i], buff[j]

		if !similarAlpha(a, b) {

			return false
		}
	}

	return true
}

func main() {

	fmt.Println("Palindrome Checker")
	res, _ := input("Input Text: ")
	fmt.Printf("Captured: %s\n", res)
	fmt.Println((func() string {

		if palindrome(res) {

			return "Palindrome"
		}

		return "Not palindrome"
	})())
}
