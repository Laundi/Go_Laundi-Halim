package main

import(
	"fmt"
)

func main() {
	var angka int
	fmt.Printf("angka: ")
	fmt Scan(&angka)

	bagi:=angka
	for i := 1; i <= bagi; i++ {
		if (angka%i == 0) {
			if i>2 {
				fmt.Print("")
			}
		}
	}
}