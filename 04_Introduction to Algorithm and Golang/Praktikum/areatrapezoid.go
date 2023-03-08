package main

import (
	"fmt"
)

func AreaTrapezoid(a int, b int, h int) int {

	return (a + b) * h / 2
}

func main() {

	fmt.Println(AreaTrapezoid(3,5,4))
}
