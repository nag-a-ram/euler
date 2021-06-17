package main

func main() {
	// using the fundamental theorem of arthimetic
	prime, num := 2, 600851475143

	for num > prime * prime {
		if num % prime == 0 {
			num /= prime
		} else { prime++ }
	}

	print("the answer is", num)
}

// written by @aliazam
