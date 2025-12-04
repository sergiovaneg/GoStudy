package main

import (
	"bufio"
	"os"

	"github.com/sergiovaneg/GoStudy/utils"
)

type Coord [2]int
type Grid map[Coord]bool

func (g Grid) getNumNeighbours(x Coord) int {
	acc := 0
	for i := x[0] - 1; i <= x[0]+1; i++ {
		for j := x[1] - 1; j <= x[1]+1; j++ {
			if g[Coord{i, j}] {
				acc++
			}
		}
	}

	return acc - 1
}

func (g Grid) recursiveRemove() int {
	acc := 0

	for x := range g {
		if g[x] && g.getNumNeighbours(x) < 4 {
			g[x] = false
			acc++
		}
	}

	if acc == 0 {
		return 0
	}
	return acc + g.recursiveRemove()
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	g := make(Grid, n*n)
	for i := 0; scanner.Scan(); i++ {
		for j, r := range scanner.Text() {
			if r == '@' {
				g[Coord{i, j}] = true
			}
		}
	}

	resA := 0
	for x := range g {
		if g.getNumNeighbours(x) < 4 {
			resA++
		}
	}
	println(resA)

	resB := g.recursiveRemove()
	println(resB)
}
