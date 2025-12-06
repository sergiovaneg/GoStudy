package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

	"github.com/sergiovaneg/GoStudy/utils"
)

func process(data []int, f func(int, int) int, carry int) int {
	for _, val := range data {
		carry = f(carry, val)
	}
	return carry
}

func parseGroupVertically(text [][]rune, startIdx int) ([]int, int) {
	vals := make([]int, 0)

	for startIdx < len(text[0]) {
		num := ""
		for _, row := range text {
			r := row[startIdx]
			if '0' <= r && '9' >= r {
				num += string(r)
			}
		}

		startIdx++
		if num == "" {
			break
		}

		val, _ := strconv.Atoi(num)
		vals = append(vals, val)
	}

	return vals, startIdx
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	var data [][]int
	text := make([][]rune, n-1)

	for i := 0; i < n-1; i++ {
		scanner.Scan()
		text[i] = []rune(scanner.Text())

		nums := regexp.MustCompile(`\d+`).FindAllString(scanner.Text(), -1)
		if data == nil {
			data = make([][]int, len(nums))
			for j := range data {
				data[j] = make([]int, n-1)
			}
		}

		for j, num := range nums {
			data[j][i], _ = strconv.Atoi(num)
		}
	}

	scanner.Scan()
	operators := regexp.MustCompile(
		`\*|\+`).FindAllString(
		scanner.Text(), -1)

	resA, resB, colIdxB := 0, 0, 0

	for colIdxA, op := range operators {
		var valsB []int
		valsB, colIdxB = parseGroupVertically(text, colIdxB)

		var f func(int, int) int
		var carry int
		switch op {
		case "+":
			f = func(a, b int) int { return a + b }
			carry = 0
		case "*":
			f = func(a, b int) int { return a * b }
			carry = 1
		}

		resA += process(data[colIdxA], f, carry)
		resB += process(valsB, f, carry)
	}

	println(resA)
	println(resB)
}
