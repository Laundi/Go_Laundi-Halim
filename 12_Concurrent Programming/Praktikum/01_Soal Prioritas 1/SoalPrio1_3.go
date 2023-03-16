package main

import (

	"fmt"
	"sync"
)

func main () {

	var wg sync.WaitGroup

	n := 12
	x := 0

	wg.Add(n)

	for i := range make([]int, n) {

		if i % 3 == 0 { x++ }
	}

	c := make(chan int, x)
	for i := range make ([]int, n) {
		go (func(i int) {
			if i % 3 == 0 {
				go (func(v int, k chan int) {
					k <- v
				})(i, c)
				go (func(v int, k chan int){
					defer wg.Done()	
					select {
						case j := <- k:
							fmt.Println(v, j)
					
					}
				} )(i,c)
			} else {
				wg.Done()
			} 
		})(i)
	}

	wg.Wait()

}