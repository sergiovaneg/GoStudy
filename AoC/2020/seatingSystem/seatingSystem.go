package main

import (
	"bufio"
	"maps"
	"os"

	"github.com/sergiovaneg/GoStudy/utils"
)

type coordinate [2]int
type grid map[coordinate]bool
type neighbourSet map[coordinate][]coordinate

func (g grid) getImmediateNeighbours() neighbourSet {
	nS := make(neighbourSet, len(g))

	for x0 := range g {
		nS[x0] = make([]coordinate, 0)
		for i := x0[0] - 1; i <= x0[0]+1; i++ {
			for j := x0[1] - 1; j <= x0[1]+1; j++ {
				x := coordinate{i, j}
				if x == x0 {
					continue
				}

				if _, ok := g[x]; ok {
					nS[x0] = append(nS[x0], x)
				}
			}
		}
	}

	return nS
}

func isValidCoord(x coordinate, n int) bool {
	if x[0] < 0 || x[1] < 0 {
		return false
	}

	if x[0] >= n || x[1] >= n {
		return false
	}

	return true
}

func offset(x, dx coordinate) coordinate {
	return coordinate{x[0] + dx[0], x[1] + dx[1]}
}

func (g grid) getClosestNeighbours(n int) neighbourSet {
	nS := make(neighbourSet, len(g))

	for x0 := range g {
		nS[x0] = make([]coordinate, 0)

		for _, dx := range [8]coordinate{
			{0, -1}, {0, 1}, {1, 0}, {-1, 0},
			{1, -1}, {1, 1}, {-1, 1}, {-1, -1},
		} {
			for x := offset(x0, dx); isValidCoord(x, n); x = offset(x, dx) {
				if _, ok := g[x]; ok {
					nS[x0] = append(nS[x0], x)
					break
				}
			}
		}
	}

	return nS
}

func (g grid) updateGrid(nS neighbourSet, tol int) grid {
	ng := make(grid, len(g))

	for x0, busy := range g {
		var nAdj int

		for _, x := range nS[x0] {
			if g[x] {
				nAdj++
			}
		}

		if busy && nAdj >= tol {
			ng[x0] = false
		} else if !busy && nAdj == 0 {
			ng[x0] = true
		} else {
			ng[x0] = busy
		}
	}

	return ng
}

func cmpGrid(a, b grid) bool {
	for x := range a {
		if a[x] != b[x] {
			return false
		}
	}

	return true
}

func (g0 grid) stabilizeGrid(nS neighbourSet, tol int) grid {
	g := maps.Clone(g0)
	for {
		ng := g.updateGrid(nS, tol)
		f := cmpGrid(g, ng)
		g = ng

		if f {
			break
		}
	}

	return g
}

func (g grid) count() int {
	var cnt int

	for _, busy := range g {
		if busy {
			cnt++
		}
	}

	return cnt
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	g0 := make(grid, n*n)

	for i := 0; scanner.Scan(); i++ {
		for j, char := range scanner.Text() {
			if char == 'L' {
				g0[coordinate{i, j}] = false
			}
		}
	}

	println(g0.stabilizeGrid(g0.getImmediateNeighbours(), 4).count())
	println(g0.stabilizeGrid(g0.getClosestNeighbours(n), 5).count())
}
