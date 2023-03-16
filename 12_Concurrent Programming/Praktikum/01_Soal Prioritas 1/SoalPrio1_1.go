package main

import (
	"fmt"
	"os"
	"time"
)

func printMultiples(x int) {
	for i := 1; ; i++ {
		fmt.Println(i * x)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	go printMultiples(2)
	time.Sleep(15 * time.Second)
	os.Exit(0)
}
