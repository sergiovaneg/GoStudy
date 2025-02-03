package main

import (
	"bufio"
	"os"
	"slices"
)

type Coordinate [2]int
type Path []Coordinate

func distance(a, b Coordinate) int {
	var d int

	for i := range 2 {
		aux := a[i] - b[i]
		if aux < 0 {
			aux = -aux
		}

		d += aux
	}

	return d
}

func sortPath(unorderedPath Path, s Coordinate) Path {
	n := len(unorderedPath)
	path := make(Path, 2, n)

	path[0] = s
	j := slices.IndexFunc(
		unorderedPath, func(x Coordinate) bool {
			return distance(x, s) == 1
		})
	path[1] = unorderedPath[j]

	for i := 2; i < n; i++ {
		j := slices.IndexFunc(unorderedPath, func(x Coordinate) bool {
			return distance(x, path[i-1]) == 1 && x != path[i-2]
		})

		path = append(path, unorderedPath[j])
	}

	return path
}

func countOptimalJumps(path Path, thr, jmp int) int {
	res, n := 0, len(path)
	for i := 0; i < n; i++ {
		for j := i + thr; j < n; j++ {
			d := distance(path[i], path[j])
			if d <= jmp && j-i-d >= thr {
				res++
			}
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

	path := make(Path, 0)
	var s Coordinate
	for i := 0; scanner.Scan(); i++ {
		for j, r := range scanner.Text() {
			if r != '#' {
				aux := Coordinate{i, j}
				path = append(path, aux)
				if r == 'S' {
					s = aux
				}
			}
		}
	}

	path = sortPath(path, s)

	println(countOptimalJumps(path, 100, 2))
	println(countOptimalJumps(path, 100, 20))
}
