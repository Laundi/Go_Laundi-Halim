package main

import (

	"fmt"
)

const (

	I = 1
	II = 2
	III = 3
	IV = 4
	V = 5
	VI = 6
	VII = 7
	VIII = 8
	IX = 9
	X = 10
	XL = 40 
	L = 50 
	XC = 90 
	C = 100 
	CD = 400
	D = 500
	CM = 900
	M = 1000
)

var maxroman = []int{M, CM, D, CD, C, XC, L, XL, X, IX, VIII, VII, VI, V, IV, III, II, I}

func findmaxroman(num int) int {

	for _, v := range maxroman {

		if v <= num {

			return v
		}
	}

	return 0
}

func romanrepr(num int) string {

	switch num {

		case I: return "I"
		case II: return "II"
		case III: return "III"
		case IV: return "IV"
		case V: return "V"
		case VI: return "VI"
		case VII: return "VII"
		case VIII: return "VIII"
		case IX: return "IX"
		case X: return "X"
		case XL: return "XL"
		case L: return "L"
		case XC: return "XC"
		case C: return "C"
		case CD: return "CD"
		case D: return "D"
		case CM: return "CM"
		case M: return "M"
	}

	return ""
}

func roman(num int) string {

	var temp string

	for {

		if num <= 0 {

			break
		}

		k := findmaxroman(num)

		temp += romanrepr(k)

		num -= k
	}

	return temp
}

func main() {

	fmt.Println(roman(4))
	fmt.Println(roman(9))
	fmt.Println(roman(23))
	fmt.Println(roman(2021))
	fmt.Println(roman(1646))
}
