package main
import "fmt"

func main() {
	line, second := []int{}, map[int]int{}

	// array of numbers 1 to 20
	for i := 1; i <= 20; i++ {
		line = append(line, i)
	}

	// prime factors of each number
	for _, mark := range line {
		divi, primary := 2, map[int]int{}

		for mark != 1 {
			if mark % divi == 0 {
				primary[divi]++
				mark /= divi
			} else { divi++ }
		}

		// storing >st occurences of factor
		for prime, time := range primary {
			if second[prime] < time {
				second[prime] = time
			}
		}
	}
		
	// getting the product of factors
	var product int = 1
	for prime, time := range second {
		for i := 0; i < time; i++ {
			product *= prime
		}
	}
	
	fmt.Printf("the answer is %d", product)
}

// made by @aliazam
