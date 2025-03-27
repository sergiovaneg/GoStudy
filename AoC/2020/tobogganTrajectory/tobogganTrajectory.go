package main

import (
	"bufio"
	"os"

	"github.com/sergiovaneg/GoStudy/utils"
)

type Grid map[[2]int]bool

func (g Grid) countCollisions(dx [2]int, height, width int) uint {
	var res uint
	var x [2]int

	for x[0] < height {
		if g[x] {
			res++
		}

		x[0] += dx[0]
		x[1] = (x[1] + dx[1]) % width
	}

	return res
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	height, _ := utils.LineCounter(file)

	grid := make(Grid)
	var width int
	for i := 0; scanner.Scan(); i++ {
		width = len(scanner.Text())
		for j, c := range scanner.Text() {
			if c == '#' {
				grid[[2]int{i, j}] = true
			}
		}
	}

	println(grid.countCollisions([2]int{1, 3}, height, width))

	resB := 1
	for _, dx := range [][2]int{
		{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}} {
		resB *= int(grid.countCollisions(dx, height, width))
	}
	println(resB)
}
