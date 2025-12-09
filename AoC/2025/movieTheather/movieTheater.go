package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

type coord [2]int

func parseCoord(line string) coord {
	var x coord
	nums := strings.Split(line, ",")
	x[0], _ = strconv.Atoi(nums[0])
	x[1], _ = strconv.Atoi(nums[1])
	return x
}

func getManhattanDistance(x coord) int {
	return utils.AbsInt(x[0]) + utils.AbsInt(x[1])
}

func getRectangleArea(x, y coord) int {
	return (utils.AbsInt(x[0]-y[0]) + 1) * (utils.AbsInt(x[1]-y[1]) + 1)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	coords := make([]coord, 0, n)

	for scanner.Scan() {
		coords = append(coords, parseCoord(scanner.Text()))
	}

	var resA int
	for i, x := range coords {
		for _, y := range coords[i+1:] {
			resA = max(resA, getRectangleArea(x, y))
		}
	}

	println(resA)
}
