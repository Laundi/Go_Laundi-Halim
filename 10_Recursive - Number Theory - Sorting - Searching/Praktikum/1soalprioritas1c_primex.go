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

func primex(n int) int {

	data := prime(n)
	return data[len(data) - 1]
}


func main() {

	fmt.Println(primex(1)) // 2
	fmt.Println(primex(5)) // 11
	fmt.Println(primex(8)) // 19
	fmt.Println(primex(9)) // 23
	fmt.Println(primex(10)) // 29
}
