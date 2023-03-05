package main

impor "fmt"

func primeNumber (Number int) bool {
	for i := 2; i*i < number, i++ {
		if number%i == 0 {
			fmt.Print(i, ",")
			fmt.Print(number, "Pembagi", number, "Bukan Prima")
			return false
		}
	}
	fmt.Print(number, "Pembagi", number, "Bilangan Prima:")
	return true
}

func main() {
	fmt.Println(primeNumber(100000007)) //true
	fmt.Println(primeNumber(13))		//true
	fmt.Println(primeNumber(17))		//true
	fmt.Println(primeNumber(20))		//false
	fmt.Println(primeNumber(35))		//false
}