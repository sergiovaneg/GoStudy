package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

type FreshList [][2]int

func isOverlap(x, y [2]int) bool {
	return x[1] >= y[0] && y[1] >= x[0]
}

func mergeRanges(x, y [2]int) [2]int {
	return [2]int{min(x[0], y[0]), max(x[1], y[1])}
}

func (l *FreshList) normalizedInsert(x [2]int) {
	mergeIdx, found := slices.BinarySearchFunc(*l, x, func(y, x [2]int) int {
		if isOverlap(x, y) {
			return 0
		}

		if x[0] > y[1] {
			return -1
		}

		return 1
	})

	if found {
		(*l)[mergeIdx] = mergeRanges((*l)[mergeIdx], x)

		currentLen := len(*l)
		for mergeIdx < currentLen-1 {
			if !isOverlap((*l)[mergeIdx], (*l)[mergeIdx+1]) {
				break
			}

			(*l)[mergeIdx] = mergeRanges((*l)[mergeIdx], (*l)[mergeIdx+1])
			*l = slices.Delete(*l, mergeIdx+1, mergeIdx+2)
			currentLen--
		}

		return
	}

	insertIdx := slices.IndexFunc(*l, func(y [2]int) bool {
		return y[0] > x[1]
	})

	if insertIdx == -1 {
		*l = append(*l, x)
	} else {
		*l = slices.Insert(*l, insertIdx, x)
	}
}

func (l FreshList) isFresh(x int) bool {
	_, found := slices.BinarySearchFunc(l, x, func(bounds [2]int, z int) int {
		if z >= bounds[0] && z <= bounds[1] {
			return 0
		}

		if z > bounds[1] {
			return -1
		}

		return 1
	})

	return found
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	l := make(FreshList, 0)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		bounds := strings.SplitN(scanner.Text(), "-", 2)
		x := [2]int{}
		x[0], _ = strconv.Atoi(bounds[0])
		x[1], _ = strconv.Atoi(bounds[1])
		l.normalizedInsert(x)
	}

	resA := 0
	for scanner.Scan() {
		target, _ := strconv.Atoi(scanner.Text())
		if l.isFresh(target) {
			resA++
		}
	}
	println(resA)

	resB := 0
	for _, bounds := range l {
		resB += bounds[1] - bounds[0] + 1
	}
	println(resB)
}
