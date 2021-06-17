package main
import ("fmt"; "strconv")

// database
var ones = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var spec = []string{"", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var tens = []string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

// drive function
func main() {
	var total int

	for i := 1; i < 1001; i++ {
		var c string = astr(i)

		switch len(c) {
		case 4:
			// 1000 only
			total += len("onethousand")
		case 3:
			// 100 to 999
			var o int = bint(c[0])
			total += len(ones[o] + "hundred")

			if c[1:] == "00" { break }
			total += len("and")
			c = c[1:]
			fallthrough
		case 1:
			// 1 to 9
			if len(c) == 1 {
				var o int = aint(c)
				total += len(ones[o])
				break
			}
			fallthrough
		case 2:
			// 10 to 99
			if c[0] == '1' && c != "10" {
				var s int = bint(c[1])
				total += len(spec[s])
			} else {
				var t int = bint(c[0])
				total += len(tens[t])

				var o int = bint(c[1])
				total += len(ones[o])
			}
			break
		}
	}
	fmt.Printf("the answer is %d\n", total)
}

// converts integer to string
func astr(i int) string {
	return strconv.Itoa(i)
}

// converts string to integer
func aint(s string) int {
	o, _ := strconv.Atoi(s)
	return o
}

// converts byte to integer
func bint(b byte) int {
	o, _ := strconv.Atoi(string(b))
	return o
}

// written by @aliazam (Alaz)
