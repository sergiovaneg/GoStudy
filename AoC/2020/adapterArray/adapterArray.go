package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"

	"github.com/sergiovaneg/GoStudy/utils"
)

func scoreAdapters(adapters []int) int {
	diffs := make(map[int]int)

	for idx, val := range adapters[1:] {
		diffs[val-adapters[idx]]++
	}

	return diffs[1] * diffs[3]
}

func recursiveCount(adapters []int) int {
	n := len(adapters)
	dp := make(map[int]int, n)
	dp[n-1] = 1

	var f func(int) int
	f = func(i int) int {
		if count, ok := dp[i]; ok {
			return count
		}

		ground := adapters[i]
		count := 0
		for j := i + 1; j < n; j++ {
			next := adapters[j]
			if next-ground > 3 {
				break
			}
			count += f(j)
		}
		dp[i] = count
		return count
	}

	return f(0)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	adapters := make([]int, n+2)

	for idx := 1; scanner.Scan(); idx++ {
		adapters[idx], _ = strconv.Atoi(scanner.Text())
	}

	slices.Sort(adapters[1 : n+1])
	adapters[0], adapters[n+1] = 0, adapters[n]+3

	println(scoreAdapters(adapters))
	println(recursiveCount(adapters))
}
