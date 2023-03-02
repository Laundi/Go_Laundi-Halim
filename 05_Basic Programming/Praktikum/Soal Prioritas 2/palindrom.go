package main

import (
	"fmt"
	"string"
)

func main() {
	teks := "radar"
	if cekKata(teks) {
		fmt.Println(teks)
		fmt.Println("Merupakan Palindrom")
	} else {
		fmt.Println(teks)
		fmt.Println("Bukan Palindrom")
	}
}

func cekKata(kata string) bool {
	kata = strings.Tolower(kata)
	for i := 0; i < len(kata)/2; i++ {
		if kata [i] ≠ kata[len(kata)-i-1] {
			return false
		} 
	}
	return true
}