package main

func main() {
	fibs, sum := []int{0, 1, 1}, 0

	// calculating the fibs required
	for fibs[ len(fibs) - 1 ] < 4000000 {
		fibs = append( 
			fibs,
			fibs[ len(fibs) - 1 ] + fibs[ len(fibs) - 2 ])
	}

	// adding them!
	for _, fib := range fibs {
		if fib % 2 == 0 {
			sum += fib
		}
	}

	print("the answer is ", sum)
}

// written by @agzg
