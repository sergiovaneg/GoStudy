package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

func intLen(x int) int {
	if x == 0 {
		return 1
	}
	return 1 + int(math.Log10(float64(x)))
}

func checkRangeA(lb, ub int) int {
	current, acc := lb, 0

	for current <= ub {
		nd := intLen(current)
		if nd%2 == 1 {
			current = utils.IPow(10, nd)
			continue
		}

		mid := utils.IPow(10, nd/2)
		if current%mid == current/mid {
			acc += current
		}

		current++
	}

	return acc
}

func checkRangeB(lb, ub int) int {
	acc := 0

	for current := lb; current <= ub; current++ {
		nd := intLen(current)

		for exp := 1; exp <= nd>>1; exp++ {
			if nd%exp != 0 {
				continue
			}

			div, aux := utils.IPow(10, exp), 0
			sub := current % div
			for range nd / exp {
				aux = aux*div + sub
			}

			if aux == current {
				acc += current
				break
			}
		}
	}

	return acc
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	accA, accB := 0, 0
	for toCheck := range strings.SplitSeq(scanner.Text(), ",") {
		bounds := strings.SplitN(toCheck, "-", 2)
		lb, _ := strconv.Atoi(bounds[0])
		ub, _ := strconv.Atoi(bounds[1])

		accA += checkRangeA(lb, ub)
		accB += checkRangeB(lb, ub)
	}

	println(accA)
	println(accB)
}
