package main

import (

	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var mu sync.Mutex

	n := 12
	x := 0

	wg.Add(n)

	for i := range make([]int, n) {
		
		if i % 3 == 0 { x++ }
	}
	
	c := make(chan int, x)

	for i := range make([]int, n) {

		go (func(i int) {

			mu.Lock()

			if i % 3 == 0 {
				
				go (func(v int, k chan int) {
					
					k <- v

				})(i,c)

				go (func(v int, k chan int) {
					
					defer wg.Done()
					
					select {
					
						case j := <- k:
							
							fmt.Println(v, j)
					}

					mu.Unlock()
				})(i, c)	
			} else {

				wg.Done()
				mu.Unlock()
			}
		})(i)
	}
	
	wg.Wait()
}
