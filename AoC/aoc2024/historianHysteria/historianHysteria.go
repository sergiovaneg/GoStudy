package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"

	"github.com/sergiovaneg/GoStudy/utils"
)

func sortedInsert[T cmp.Ordered](ts []T, t T) []T {
	i, _ := slices.BinarySearch(ts, t)
	return slices.Insert(ts, i, t)
}

func getValues(line string) [2]int {
	matches := regexp.MustCompile("[0-9]+").FindAllString(line, 2)
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
	for idx := 0; idx < n; idx++ {
		a := listA[idx]
		b := listB[idx]
		if a > b {
			res0 += a - b
		} else {
			res0 += b - a
		}
	}

	fmt.Println(res0)

	res1 := 0
	for i, j := 0, 0; i < n && j < n; {
		for j < n && listA[i] > listB[j] {
			j++
		}
		if j == n {
			break
		}
		for i < n && listA[i] < listB[j] {
			i++
		}
		if i == n {
			break
		}

		if listA[i] != listB[j] {
			continue
		}

		ref, na, nb := listA[i], 0, 0
		for i < n && listA[i] == ref {
			na++
			i++
		}
		for j < n && listB[j] == ref {
			nb++
			j++
		}

		res1 += ref * na * nb
	}

	fmt.Println(res1)
}
