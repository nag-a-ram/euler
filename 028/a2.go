// Copyright 2021 Alee Azam
// --------- ---- ---- ----
//
// real    0m0.184s
// user    0m0.217s
// sys     0m0.065s

// Starting with 1, the loop adds each "corner" digit of the diagonals to
// total.
// 
// For each spiral, if the width is w then the successive digit is w+2.
//
// This means that the distance between the two corner digits (n) also
// increases by 2 in each successive spiral (n+2).

package main

import "fmt"

func main() {
	n, w, total := 2, 3, 0
	
	for i := 1; w <= 1001; i += n {		
		total += i

		if i == (w*w) {
			n += 2
			w += 2
		}
	}
	fmt.Println(total)
}
