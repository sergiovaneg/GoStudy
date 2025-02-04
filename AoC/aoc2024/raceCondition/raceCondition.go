package main

import (
	"bufio"
	"os"
	"slices"
)

const THR = 100

type Coordinate [2]int
type Chain [2][2]int
type Path []Coordinate

func distance(a, b Coordinate) int {
	var d int

	for i := range 2 {
		if a[i] > b[i] {
			d += a[i] - b[i]
		} else {
			d += b[i] - a[i]
		}
	}

	return d
}

func (prev Chain) pathIndexer(x Coordinate) bool {
	return distance(x, prev[1]) == 1 && x != prev[0]
}

func sortPath(unorderedPath Path, start Coordinate) Path {
	n := len(unorderedPath)
	path := make(Path, n)

	path[0] = start
	prev := Chain{start, start}

	for i := 1; i < n; i++ {
		j := slices.IndexFunc(unorderedPath, prev.pathIndexer)
		path[i] = unorderedPath[j]

		prev[0] = prev[1]
		prev[1] = path[i]
	}

	return path
}

func (path Path) countOptimalJumps(thr, jmp int) int {
	n := len(path)
	ub := n - thr

	c := make(chan int, ub)
	defer close(c)

	spawnable := func(lb int) {
		res := 0
		for j := lb + thr; j < n; j++ {
			d := distance(path[lb], path[j])
			if d <= jmp && j-lb-d >= thr {
				res++
			}
		}
		c <- res
	}

	for i := range ub {
		go spawnable(i)
	}

	res := 0
	for range ub {
		res += <-c
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

	path := make(Path, 0)
	var s Coordinate
	for i := 0; scanner.Scan(); i++ {
		for j, r := range scanner.Text() {
			if r == '#' {
				continue
			}

			aux := Coordinate{i, j}
			path = append(path, aux)
			if r == 'S' {
				s = aux
			}
		}
	}

	path = sortPath(path, s)

	println(path.countOptimalJumps(THR, 2))  // Part A
	println(path.countOptimalJumps(THR, 20)) // Part B
}
