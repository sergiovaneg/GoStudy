package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

type coord [3]int

const connectionNumber = 1000

func parseCoord(line string) coord {
	var x coord

	for i, num := range strings.Split(line, ",") {
		val, _ := strconv.Atoi(num)
		x[i] = val
	}

	return x
}

func d2(x, y coord) int {
	d := 0

	for idx := range 3 {
		d += (x[idx] - y[idx]) * (x[idx] - y[idx])
	}

	return d
}

func getDistanceMatrix(arr []coord) [][]int {
	ret := make([][]int, len(arr))

	for i, x := range arr {
		ret[i] = make([]int, len(arr))
		for j, y := range arr {
			if i == j {
				ret[i][j] = -1
			} else {
				ret[i][j] = d2(x, y)
			}
		}
	}

	return ret
}

func cmpDistance(a, b int) int {
	if a == -1 {
		return 1
	} else if b == -1 {
		return -1
	} else {
		return a - b
	}
}

func singleIter(
	arr []coord,
	dMatrix *[][]int,
	membership *[][]coord,
	minDVec, minDIdx *[]int) (coord, coord) {
	// Update nearest-box slices
	for i, dRow := range *dMatrix {
		if (*minDVec)[i] > 0 {
			continue
		}

		(*minDVec)[i] = slices.MinFunc(dRow, cmpDistance)
		(*minDIdx)[i] = slices.Index(dRow, (*minDVec)[i])
	}

	// Select nearest pair
	bestI := slices.Index(*minDVec, slices.Min(*minDVec))
	bestJ := (*minDIdx)[bestI]
	boxI := arr[bestI]
	boxJ := arr[bestJ]

	// Mask distances
	(*dMatrix)[bestI][bestJ] = -1
	(*dMatrix)[bestJ][bestI] = -1
	(*minDVec)[bestI] = 0
	(*minDVec)[bestJ] = 0

	// Locate pre-assigned boxes
	idI, idJ := -1, -1
	for id, members := range *membership {
		for _, x := range members {
			if boxI == x {
				idI = id
			}
			if boxJ == x {
				idJ = id
			}
		}
	}

	if idI > -1 && idJ > -1 {
		if idI != idJ {
			(*membership)[idI] = append((*membership)[idI], (*membership)[idJ]...)
			*membership = slices.Delete(*membership, idJ, idJ+1)
		}
	} else if idI > -1 {
		(*membership)[idI] = append((*membership)[idI], boxJ)
	} else if idJ > -1 {
		(*membership)[idJ] = append((*membership)[idJ], boxI)
	} else {
		*membership = append(*membership, []coord{boxI, boxJ})
	}

	return boxI, boxJ
}

func iterativeConnect(arr []coord) (int, int) {
	dMatrix := getDistanceMatrix(arr)
	membership := make([][]coord, 0)

	minDVec := make([]int, len(arr))
	minDIdx := make([]int, len(arr))
	for range connectionNumber {
		singleIter(arr, &dMatrix, &membership, &minDVec, &minDIdx)
	}

	setSizes := make([]int, 0, len(membership))
	for _, set := range membership {
		setSizes = append(setSizes, len(set))
	}
	slices.SortFunc(setSizes, func(a, b int) int { return b - a })
	setSizes = setSizes[:3]

	var x, y coord
	for len(membership) > 1 && len(membership[0]) < len(arr) {
		x, y = singleIter(arr, &dMatrix, &membership, &minDVec, &minDIdx)
	}

	return setSizes[0] * setSizes[1] * setSizes[2], x[0] * y[0]
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	arr := make([]coord, 0, n)

	for scanner.Scan() {
		arr = append(arr, parseCoord(scanner.Text()))
	}

	println(iterativeConnect(arr))
}
