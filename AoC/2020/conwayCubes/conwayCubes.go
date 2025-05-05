package main

import (
	"bufio"
	"maps"
	"os"
	"slices"
)

const dim = 4

type Coordinate [dim]int
type Grid map[Coordinate]bool

func add(x1, x2 Coordinate) Coordinate {
	var x Coordinate

	for idx := range dim {
		x[idx] = x1[idx] + x2[idx]
	}

	return x
}

func (x0 Coordinate) getNeighbours() []Coordinate {
	var f func(int) []Coordinate
	f = func(i int) []Coordinate {
		if i == dim {
			return []Coordinate{x0}
		}

		nextLevel := f(i + 1)
		neighbourhood := make([]Coordinate, 0, 3*len(nextLevel))

		var x1 Coordinate
		for _, x2 := range nextLevel {
			for _, dx := range [3]int{-1, 0, 1} {
				x1[i] = dx
				neighbourhood = append(neighbourhood, add(x1, x2))
			}
		}

		return neighbourhood
	}

	return slices.DeleteFunc(
		f(0),
		func(x Coordinate) bool { return x == x0 })
}

func (g *Grid) pad() {
	newG := make(Grid)

	for x0, active := range *g {
		if !active {
			continue
		}

		newG[x0] = true
		for _, x := range x0.getNeighbours() {
			if _, ok := newG[x]; !ok {
				newG[x] = false
			}
		}
	}

	*g = newG
}

func (g Grid) getNActiveNeighbours(x0 Coordinate) int {
	var res int
	for _, x := range x0.getNeighbours() {
		if g[x] {
			res++
		}
	}
	return res
}

func (g *Grid) evolve() {
	g.pad()
	newG := maps.Clone(*g)

	for x0, active := range *g {
		n := g.getNActiveNeighbours(x0)
		if active {
			newG[x0] = n == 2 || n == 3
		} else {
			newG[x0] = n == 3
		}
	}

	*g = newG
}

func (g Grid) score() int {
	var res int
	for _, active := range g {
		if active {
			res++
		}
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
	g := make(Grid)

	for i := 0; scanner.Scan(); i++ {
		for j, c := range scanner.Text() {
			g[Coordinate{i, j}] = c == '#'
		}
	}

	for range 6 {
		g.evolve()
	}
	println(g.score())
}
