package main
import ("strconv"; "fmt")

func main() {
	var best int

	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			product := i * j
			str := strconv.Itoa(product)
			rev := reverse(str)

			if rev == str && product > best{
				best = product
			}
		}
	}

	fmt.Printf("the answer is %d", best)
}

func reverse(s string) string {
	var out string
	for _, char := range s {
		out = string(char) + out
	}
	return out
}

// written by @aliazam
