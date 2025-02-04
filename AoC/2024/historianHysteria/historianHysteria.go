package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

func sortedInsert[T cmp.Ordered](ts []T, t T) []T {
	i, _ := slices.BinarySearch(ts, t)
	return slices.Insert(ts, i, t)
}

func getValues(line string) [2]int {
	matches := strings.Split(line, "   ")
	a, _ := strconv.Atoi(matches[0])
	b, _ := strconv.Atoi(matches[1])
	return [2]int{a, b}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	listA, listB := make([]int, 0, n), make([]int, 0, n)
	for scanner.Scan() {
		nums := getValues(scanner.Text())
		listA = sortedInsert(listA, nums[0])
		listB = sortedInsert(listB, nums[1])
	}

	res0 := 0
	mA, mB := make(map[int]int, n), make(map[int]int, n)
	for idx := 0; idx < n; idx++ {
		a := listA[idx]
		b := listB[idx]

		mA[a]++
		mB[b]++

		if a > b {
			res0 += a - b
		} else {
			res0 += b - a
		}
	}

	fmt.Println(res0)

	res1 := 0
	for k, v := range mA {
		res1 += k * v * mB[k]
	}

	fmt.Println(res1)
}
