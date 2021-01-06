package solution

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	out := []string{}
	var s string

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			s = "FizzBuzz"
		} else if i%5 == 0 {
			s = "Buzz"
		} else if i%3 == 0 {
			s = "Fizz"
		} else {
			s = strconv.Itoa(i)
		}

		out = append(out, s)
	}

	return out
}
