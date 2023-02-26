package main

import (
	"fmt"
	"strconv"
)

type fraction struct {
	n, d float64
}

func main() {
	var fractions = []fraction{}

	// considering the fraction: i/j
	for i := 10.00; i < 100.00; i++ {
		for j := 10.00; j < 100.00; j++ {			
			// filtering "less than one in value",
			// and skipping "trivial" cases
			if (i >= j) || (indexI(i, 2) == 0 && indexI(j, 2) == 0) { continue }

			if (i/j == indexI(i, 1)/indexI(j, 1) && indexI(i, 2) == indexI(j, 2)) ||
			   (i/j == indexI(i, 2)/indexI(j, 2) && indexI(i, 1) == indexI(j, 1)) ||
			   (i/j == indexI(i, 1)/indexI(j, 2) && indexI(i, 2) == indexI(j, 1)) ||
			   (i/j == indexI(i, 2)/indexI(j, 1) && indexI(i, 1) == indexI(j, 2)) {
					fractions = append(fractions, fraction{i, j})
			}
		}
	}
	var n, d = 1.00, 1.00
	for _, f := range fractions {
		n *= f.n
		d *= f.d
	}
	// 0.01 = 1/100
	fmt.Println(n/d)
}

// indexI returns the float at (i)ndex of
// the provided (f)loat.
func indexI(f float64, i int) float64 {
	s := strconv.Itoa(int(f))
	c, _ := strconv.Atoi(string(s[i - 1]))

	return float64(c)
}

// written by @agzg
