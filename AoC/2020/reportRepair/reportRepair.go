package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"

	"github.com/sergiovaneg/GoStudy/utils"
)

const target = 2020

func recSearch(entries []int, rem int, n int) int {
	if n == 1 {
		if idx, found := slices.BinarySearch(entries, rem); found {
			return entries[idx]
		}
		return -1
	}

	for i, v1 := range entries {
		if v1 > rem {
			break
		}

		if v2 := recSearch(entries[i+1:], rem-v1, n-1); v2 != -1 {
			return v1 * v2
		}
	}

	return -1
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	entries := make([]int, n)

	for i := 0; scanner.Scan(); i++ {
		entries[i], _ = strconv.Atoi(scanner.Text())
	}

	slices.Sort(entries)
	println(recSearch(entries, target, 2))
	println(recSearch(entries, target, 3))
}
