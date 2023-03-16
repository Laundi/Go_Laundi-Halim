package main

import (
	"fmt"
)

func main() {
	m := make(chan int) // buat channel untuk mengirim bilangan kelipatan 3

	go kelipatan3(m) // jalankan go routine untuk mengirim nilai ke channel

	for n := range m { // baca nilai dari channel dan cetak satu per satu
		fmt.Println(n)
	}
}

func kelipatan3(m chan int) {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			m <- i // kirim bilangan kelipatan 3 ke channel
		}
	}
	close(m) // tutup channel setelah selesai mengirim nilai
}
