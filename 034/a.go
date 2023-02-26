// Written by @agzg

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// concludes that our limit is 2540160 (or 7 * 9!)
	//
	// i.e. numbers greater than this cannot fulfill the requirement of the
	// question, "equal to the sum of the factorial of their digits"

	var limit, answer int

	for i := 9; limit == 0; i = i * 10 + 9 {
		accum, length := 0, len(strconv.Itoa(i))

		for j := 0; j < length; j++ {
			accum += fact(9)
		}
		if i > accum {
			limit = length * fact(9)
		}
	}
	// fmt.Println("limit", limit)

	for i := 3; i <= limit; i++ {
		var accum int

		for _, c := range strconv.Itoa(i) {
			n, _ := strconv.Atoi(string(c))
			accum += fact(n)
		}
		if accum == i {
			answer += i
		}
	}
	fmt.Println(answer)
}

func fact(n int) int {
	/*
	if n <= 2 {
		return n
	}
	*/

	o := 1
	for i := 2; i <= n; i++ {
		o *= i
	}
	return o
}
