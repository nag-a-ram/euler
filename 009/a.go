package main
import "fmt"

func main() {
	for a := 1; a <= 999; a++ {
		for b := 1; b <= 999; b++ {
			c := 1000 - a - b
			
			if (pow(a, 2) + pow(b, 2)) == pow(c, 2) {
				fmt.Printf("the answer is %d\n", a*b*c)
				euclid()
				return
			}
		}
	}
}

// another cool method
func euclid() {
	/* 

		* Using Euclid's Formula *
	
			 for m > n > 0
	a = m2 - n2; b = 2mn; c = m2 + n2

	*/

	for m := 1; m <= 999; m++ {
		for n := 1; n <= 999; n++ {
			ssum := 2*m*(m + n)

			if (ssum == 1000) && (m > n) {
				answer := 2*m*n*(pow(m, 4) - pow(n, 4))
				fmt.Printf("\neuclid's\nthe answer is %d\n", 
					answer)
				return			
			}
		}
	}
}

func pow(n, p int) int {
	var out int = 1

	for i := 1; i <= p; i++ {
		out *= n
	}

	return out
}

// written by @aliazam
