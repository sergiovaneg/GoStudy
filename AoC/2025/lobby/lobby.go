package main

import (
	"bufio"
	"os"
	"slices"
)

func recursive_select(arr []rune, missing int, acc int) int {
	if missing == 0 {
		return acc
	}
	n := len(arr)

	max_val := slices.Max(arr[:n+1-missing])
	max_idx := slices.Index(arr[:n+1-missing], max_val)

	return recursive_select(
		arr[max_idx+1:], missing-1, 10*acc+int(max_val-'0'),
	)
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
		accA += recursive_select(arr, 2, 0)
		accB += recursive_select(arr, 12, 0)
	}
	println(accA)
	println(accB)
}
