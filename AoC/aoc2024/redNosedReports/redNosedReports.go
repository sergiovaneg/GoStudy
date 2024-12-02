package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

func parseLevel(line string) []int {
	nums := strings.Split(line, " ")
	levels := make([]int, len(nums))
	for idx, num := range nums {
		if aux, err := strconv.Atoi(num); err == nil {
			levels[idx] = aux
		}
	}
	return levels
}

func isSafeLevel(level []int, tol bool) bool {
	dir := 0
	if level[1]-level[0] > 0 {
		dir = 1
	} else {
		dir = -1
	}

	for n, idx := len(level), 1; idx < n; idx++ {
		criterion := dir * (level[idx] - level[idx-1])
		if criterion < 1 || criterion > 3 {
			if !tol {
				return false
			}

			aux_0 := make([]int, n-1)
			copy(aux_0[:idx-1], level[:idx-1])
			copy(aux_0[idx-1:], level[idx:])

			aux_1 := make([]int, n-1)
			copy(aux_1[:idx], level[:idx])
			copy(aux_1[idx:], level[idx+1:])

			return isSafeLevel(aux_0, false) || isSafeLevel(aux_1, false)
		}
	}

	return true
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	levels := make([][]int, n)
	for idx := 0; scanner.Scan(); idx++ {
		levels[idx] = parseLevel(scanner.Text())
	}

	res0 := 0
	for _, level := range levels {
		if isSafeLevel(level, false) {
			res0++
		}
	}

	fmt.Println(res0)

	res1 := 0
	for _, level := range levels {
		if isSafeLevel(level, true) {
			res1++
		}
	}

	fmt.Println(res1)
}
