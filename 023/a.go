// Written by @agzg
// run 10.8811991s, build 10.740976852s

package main

import "fmt"

func main() {
	abundant, answer := []int{}, 0

	for i := 1; i < 28123; i++ {
		for j, a := range abundant {
			accum := i - a

			for _, b := range abundant[j:] {
				if b == accum {
					goto abundance
				}
			}
		}
		answer += i
		
		abundance:
			accum := 1

			for k := 2; k < i; k++ {
				if i % k == 0 {
					accum += k
				}
			}
			if accum > i {
				abundant = append(abundant, i)
			}
	}
	fmt.Println(answer)
}
