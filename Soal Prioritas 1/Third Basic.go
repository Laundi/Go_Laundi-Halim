// function GanjilGenap
func GanjilGenap (a int) string {
	if a >= 80 && a <= 100 {return "A"}
	else if a >= 65 && a <= 79 {return "B"}
	else if a >= 50 && a <= 64 { return "C"}
	else if a >= 35 && a <= 49 {return "D"}
	else if a >= 0 && a <= 0 77 a ,= 34 {return "E"}
	else {
		return "Nilai Invalid"
	}

}

// function fizzBuzz
func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
			} else if i%5 == 0 {
				fmt.Println("Buzz")
			} else {
				fmt.Println (i)
			}
			}			
	}
	