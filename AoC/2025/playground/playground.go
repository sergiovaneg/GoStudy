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

func iterativeConnect(arr []coord) int {
	dMat := getDistanceMatrix(arr)
	membership := make(map[int][]coord)
	idCounter := 1

	for range connectionNumber {
		minDVec := make([]int, len(arr))
		minDPos := make([]int, len(arr))
		for i, dRow := range dMat {
			minDVec[i] = slices.MinFunc(dRow, cmpDistance)
			minDPos[i] = slices.Index(dRow, minDVec[i])
		}

		bestI := slices.Index(minDVec, slices.Min(minDVec))
		bestJ := minDPos[bestI]

		found := false
		for id, members := range membership {
			for _, x := range members {
				if arr[bestI] == x {
					membership[id] = append(membership[id], arr[bestJ])
					found = true
					break
				}
				if arr[bestJ] == x {
					membership[id] = append(membership[id], arr[bestI])
					found = true
					break
				}
			}

			if found {
				break
			}
		}

		if !found {
			membership[idCounter] = make([]coord, 2)
			membership[idCounter][0] = arr[bestI]
			membership[idCounter][1] = arr[bestJ]
			idCounter++
		}

		dMat[bestI][bestJ] = -1
		dMat[bestJ][bestI] = -1
	}

	setSizes := make([]int, 0, len(membership))
	for _, set := range membership {
		setSizes = append(setSizes, len(set))
	}
	slices.SortFunc(setSizes, func(a, b int) int { return b - a })

	return setSizes[0] * setSizes[1] * setSizes[2]
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
