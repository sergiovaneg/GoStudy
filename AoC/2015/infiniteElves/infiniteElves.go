package main

import (
	"math"
	"slices"
)

const INPUT = 36000000

func getFactors(n int) []int {
	if n == 1 {
		return []int{1}
	}
	result := make([]int, 0, n>>1+1)

	limit := int(math.Sqrt(float64(n))) + 1
	for i := 1; i < limit; i++ {
		if n%i == 0 {
			result = append(result, i)
			if j := n / i; j != i {
				result = append(result, j)
			}
		}
	}

	return result
}

func getPresentsA(n int) int {
	res := 0
	for _, val := range getFactors(n) {
		res += val
	}
	return res * 10
}

func getPresentsB(n int) int {
	res := 0
	factors := getFactors(n)
	slices.Sort(factors)
	preserveIdx := slices.IndexFunc(factors, func(x int) bool {
		return n/x <= 50
	})
	factors = factors[preserveIdx:]
	for _, val := range factors {
		res += val
	}
	return res * 11
}

func main() {
	for x := 1; ; x++ {
		if getPresentsA(x) > INPUT {
			println(x)
			break
		}
	}
	for x := 1; ; x++ {
		if getPresentsB(x) > INPUT {
			println(x)
			break
		}
	}
}
