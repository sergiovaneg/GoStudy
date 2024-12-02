package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

func isSafeLevel(level []int, tol int) bool {
	dir, n := 0, len(level)
	if n < 3 {
		return true
	}

	if level[1] > level[0] {
		dir = 1
	} else {
		dir = -1
	}

	for idx := 1; idx < n; idx++ {
		criterion := dir * (level[idx] - level[idx-1])
		if criterion < 1 || criterion > 3 {
			if tol == 0 {
				return false
			}

			aux_0, aux_1 := slices.Clone(level), slices.Clone(level)
			aux_0 = slices.Delete(aux_0, idx-1, idx)
			aux_1 = slices.Delete(aux_1, idx, idx+1)

			return isSafeLevel(aux_0, tol-1) || isSafeLevel(aux_1, tol-1)
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
		if isSafeLevel(level, 0) {
			res0++
		}
	}

	fmt.Println(res0)

	res1 := 0
	for _, level := range levels {
		if isSafeLevel(level, 1) {
			res1++
		}
	}

	fmt.Println(res1)
}
