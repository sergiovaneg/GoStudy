package main

import (
	"bufio"
	"os"
	"slices"
)

func recursive_select(arr []rune, missing int) []rune {
	if missing == 0 {
		return []rune{}
	}
	n := len(arr)

	max_val := slices.Max(arr[:n+1-missing])
	max_idx := slices.Index(arr[:n+1-missing], max_val)

	return append(
		[]rune{max_val},
		recursive_select(arr[max_idx+1:], missing-1)...)
}

func get_joltage(arr []rune) int {
	val := 0

	for _, r := range arr {
		val = val*10 + int(r-'0')
	}

	return val
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// n, _ := utils.LineCounter(file)

	accA, accB := 0, 0
	for scanner.Scan() {
		arr := []rune(scanner.Text())
		accA += get_joltage(recursive_select(arr, 2))
		accB += get_joltage(recursive_select(arr, 12))
	}
	println(accA)
	println(accB)
}
