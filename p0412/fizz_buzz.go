package p0412

import "fmt"

func FizzBuzz(n int) []string {
	res := make([]string, n)
	counter3, counter5 := 0, 0

	for idx := 1; idx <= n; idx++ {
		counter3++
		counter5++
		if counter3 == 3 && counter5 == 5 {
			counter3 = 0
			counter5 = 0
			res[idx-1] = "FizzBuzz"
		} else if counter3 == 3 {
			counter3 = 0
			res[idx-1] = "Fizz"
		} else if counter5 == 5 {
			counter5 = 0
			res[idx-1] = "Buzz"
		} else {
			res[idx-1] = fmt.Sprint(idx)
		}
	}

	return res
}
