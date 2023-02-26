// Written by @agzg
// r: 670ms; x: 650ms

package main

import (
	"fmt"
)

// fail: reverse Collatzing
// t&er: initial odds yield longer seqs
func main() {
	b, l := 0, 0

	for i := 1; i < 1000000; i += 2 {
		lc := len(collatz(i))
		if lc > l {
			b = i
			l = lc
		}
	}
	fmt.Println(b)
}

// collatz constructs and returns a collatz sequence
// that starts with n
func collatz(n int) (s []int) {
	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		s = append(s, n)
	}
	return
}
