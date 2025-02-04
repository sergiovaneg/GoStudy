package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

func processLine(line string, c chan<- [2]int) {
	nums := strings.Split(line, "\t")
	n := len(nums)
	vals := make([]int, n)

	for i, num := range nums {
		val, _ := strconv.Atoi(num)
		vals[i] = val
	}

	slices.Sort(vals)

	quotient := -1
	for i, q := range vals {
		for _, p := range vals[i+1:] {
			if p%q == 0 {
				quotient = p / q
			}
		}

		if quotient != -1 {
			break
		}
	}

	c <- [2]int{vals[n-1] - vals[0], quotient}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	c := make(chan [2]int, n)

	for scanner.Scan() {
		go processLine(scanner.Text(), c)
	}

	var resA, resB int
	for range n {
		aux := <-c
		resA += aux[0]
		resB += aux[1]
	}

	println(resA)
	println(resB)
}
