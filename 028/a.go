package main
import "fmt"

func main() {
	var n int = 2	// add-i
	var t int = 3	// sq-spiral width
	var s int = 0	// sum
	
	for i := 1; t != 1003; i += n {
		// fmt.Printf("%d-%d-%d\n", i, n, t)
		s += i

		if i == (t*t) {
			n += 2
			t += 2
		}
	}

	fmt.Println("the solution is", s)
}

// written by @aliazam (Alaz)
