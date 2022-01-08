package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var sum int

	for j := 123456789; j < 10000000000; j++ {
		a := strconv.Itoa(j)
		a = strings.Repeat("0", 10-len(a)) + a

		if strings.Contains(a, "0") && strings.Contains(a, "1") && strings.Contains(a, "2") && strings.Contains(a, "9") &&
			strings.Contains(a, "3") && strings.Contains(a, "4") && strings.Contains(a, "5") &&
			strings.Contains(a, "6") && strings.Contains(a, "7") && strings.Contains(a, "8") {

			// d2d3d4 should be divisible by 2.
			i, err := strconv.Atoi(a[1:5])
			if err != nil {
				panic(err)
			}

			if i%2 != 0 {
				continue
			}

			// d3d4d5 should be divisible by 3.
			i, err = strconv.Atoi(a[2:6])
			if err != nil {
				panic(err)
			}

			if i%3 != 0 {
				continue
			}

			// d4d5d6 should be divisible by 5.
			i, err = strconv.Atoi(a[3:6])
			if err != nil {
				panic(err)
			}

			if i%5 != 0 {
				continue
			}

			// d5d6d7 should be divisible by 7.
			i, err = strconv.Atoi(a[5:7])
			if err != nil {
				panic(err)
			}

			if i%7 != 0 {
				continue
			}

			// d6d7d8 should be divisible by 11.
			i, err = strconv.Atoi(a[6:8])
			if err != nil {
				panic(err)
			}

			if i%11 != 0 {
				continue
			}

			// d7d8d9 should be divisible by 13.
			i, err = strconv.Atoi(a[7:9])
			if err != nil {
				panic(err)
			}

			if i%13 != 0 {
				continue
			}

			// d8d9d10 should be divisible by 17.
			i, err = strconv.Atoi(a[8:10])
			if err != nil {
				panic(err)
			}

			if i%17 != 0 {
				continue
			}

			sum += j
		}
	}
	fmt.Println(sum)
}
