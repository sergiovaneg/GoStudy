package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"

	"github.com/sergiovaneg/GoStudy/utils"
)

const BUFFERSIZE = 25

func getInvalid(srcList []int, m int) int {
	buffer := make([]int, 0, m)

	for _, num := range srcList[:m] {
		buffer, _ = utils.SortedUniqueInsert(buffer, num)
	}

	for idx, num := range srcList[m:] {
		var found bool

		for i, a := range buffer {
			_, f := slices.BinarySearch(buffer[i+1:], num-a)
			if f {
				found = true
				break
			}
		}

		if !found {
			return num
		}

		delIdx, _ := slices.BinarySearch(buffer, srcList[idx])
		buffer = slices.Delete(buffer, delIdx, delIdx+1)
		buffer, _ = utils.SortedUniqueInsert(buffer, num)
	}

	return -1
}

func findWeakness(srcList []int, target int) int {
	lb, ub := 0, 2
	acc := srcList[0] + srcList[1]

	for acc != target {
		if acc > target {
			if ub-lb == 2 {
				acc += srcList[ub] - srcList[lb]
				lb++
				ub++
			} else {
				acc -= srcList[lb]
				lb++
			}
		} else {
			acc += srcList[ub]
			ub++
		}
	}

	return slices.Min(srcList[lb:ub]) + slices.Max(srcList[lb:ub])
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	numList := make([]int, n)

	for idx := 0; scanner.Scan(); idx++ {
		numList[idx], _ = strconv.Atoi(scanner.Text())
	}

	invalid := getInvalid(numList, BUFFERSIZE)
	println(invalid)
	println(findWeakness(numList, invalid))
}
