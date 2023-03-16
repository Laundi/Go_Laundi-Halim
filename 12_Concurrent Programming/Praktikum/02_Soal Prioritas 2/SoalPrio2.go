package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	c := make(chan map[rune]int)
	var wg sync.WaitGroup

	var count = func(text string) {
		defer wg.Done()
		tempMap := make(map[rune]int)
		for _, t := range text {
			tempMap[t]++
		}
		c <- tempMap
	}

text := "Loream ipsum dolor sit amet, consectetur elit, sed do eiusmod tempor incididunt ut labore et dolore magnaaliqua"
textSplit := strings.Split(text, " ")
for _, tx := range textSplit {
	wg.Add(1)
	go count(tx)
}

	tempMap := make(map[rune]int)
	for i := 0; i < len(textSplit); i++ {
		temp := <- c
		for k, v := range temp {
			tempMap[k] += v
		}
	}

	for k, v := range tempMap {
		fmt.Printf("%c : %d\n", k, v)
	}
}
