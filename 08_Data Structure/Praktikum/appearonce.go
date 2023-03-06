package main

import (
	"fmt"
	"strconv"
)

func appearOnce(context string) []int {

	data := make(map[int]int)

	for _, v := range []byte(context) {

		num, err := strconv.ParseInt(string(v), 0, 32)

		if err != nil {

			panic(err)
			continue
		}
		data[int(num)]++
	}

	result := []int{}

	for i, v := range data {

		if v == 1 {

			result = append(result, i)
		}
	}

	return result
}

func main() {

	fmt.Println(appearOnce("1234123"))
	fmt.Println(appearOnce("76523752"))
	fmt.Println(appearOnce("12345"))
	fmt.Println(appearOnce("1122334455"))
	fmt.Println(appearOnce("0872504"))
}