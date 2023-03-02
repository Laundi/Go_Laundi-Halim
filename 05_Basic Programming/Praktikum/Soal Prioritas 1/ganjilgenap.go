cpackage main

import "fmt"

func main () {
	var angka int = 1

	fmt.Print("Ketik Angka = ")
	fmt.Scanln(&angka)

	if angka%2 != 0 {
		fmt.Println(angka, "Merupakan Bilangan Ganjil")
	} else {
		fmt.Println(angka, "Merupakan Bilangan Genap")
	}
}