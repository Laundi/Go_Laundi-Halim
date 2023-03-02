package main

import "fmt"

func main () {
	
	var b, t, a float64

	fmt.Print("Ketik Nilai Sisi A = ")
	fmt.Scanln(&a)
	fmt.Print("Ketik Nilai Sisi B = ")
	fmt.Scanln(&b)
	fmt.Print("Ketik Nilai Tinggi = ")
	fmt.Scanln(&t)

	var hasil float64 = (1.0 / 2) * ((a+b)*t)

	if hasil != 0 {
		fmt.Println("Hasil Nilai Trapesium Adalah", hasil)
	} else {
		fmt.Println("Masukkan Nilai")
	}

}