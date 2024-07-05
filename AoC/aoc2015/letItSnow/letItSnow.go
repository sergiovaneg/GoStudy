package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

const FirstCode = 20151125
const Factor = 252533
const Divisor = 33554393

func getNextPosition(i, j int) (int, int) {
	nextI, nextJ := i-1, j+1
	if nextI == 0 {
		nextI, nextJ = nextJ, 1
	}
	return nextI, nextJ
}

func getCode(row, col int) int {
	i, j := 1, 1
	code := FirstCode
	for !(i == row && j == col) {
		i, j = getNextPosition(i, j)
		code = (code * Factor) % Divisor
	}
	return code
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var target [2]int
	for idx, num := range regexp.MustCompile("([0-9]+)").FindAllString(scanner.Text(), 2) {
		val, _ := strconv.Atoi(num)
		target[idx] = val
	}

	println(getCode(target[0], target[1]))
}
