package main

import (
	"bufio"
	"maps"
	"os"

	"github.com/sergiovaneg/GoStudy/utils"
)

type grid map[[2]int]bool
type counter func(grid, [2]int) int

func countAdjacent(g grid, x [2]int) int {
	var cnt int

	for i := x[0] - 1; i <= x[0]+1; i++ {
		for j := x[1] - 1; j <= x[1]+1; j++ {
			if g[[2]int{i, j}] {
				cnt++
			}
		}
	}
	if g[x] {
		cnt--
	}

	return cnt
}

func isValidCoordinate(x [2]int, n int) bool {
	if x[0] < 0 || x[1] < 0 {
		return false
	}

	if x[0] >= n || x[1] >= n {
		return false
	}

	return true
}

func offset(x, dx [2]int) [2]int {
	return [2]int{x[0] + dx[0], x[1] + dx[1]}
}

func countNearest(g grid, x0 [2]int, n int) int {
	var cnt int

	for _, dx := range [8][2]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	} {
		var f bool
		for x := offset(x0, dx); isValidCoordinate(x, n); x = offset(x, dx) {
			if v, ok := g[x]; ok {
				f = v
				break
			}
		}
		if f {
			cnt++
		}
	}

	return cnt
}

func (g grid) updateGrid(rule counter, tol int) grid {
	newGrid := make(grid, len(g))

	for k, v := range g {
		nAdj := rule(g, k)
		if v && nAdj >= tol {
			newGrid[k] = false
		} else if !v && nAdj == 0 {
			newGrid[k] = true
		} else {
			newGrid[k] = v
		}
	}

	return newGrid
}

func cmpGrid(a, b grid) bool {
	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}

	return true
}

func (g grid) count() int {
	var cnt int

	for _, v := range g {
		if v {
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
				g0[[2]int{i, j}] = false
			}
		}
	}

	g := maps.Clone(g0)
	for {
		ng := g.updateGrid(countAdjacent, 4)
		f := cmpGrid(g, ng)
		g = ng

		if f {
			break
		}
	}
	println(g.count())

	g = maps.Clone(g0)
	for {
		ng := g.updateGrid(func(g grid, i [2]int) int {
			return countNearest(g, i, n)
		}, 5)
		f := cmpGrid(g, ng)
		g = ng

		if f {
			break
		}
	}
	println(g.count())
}
