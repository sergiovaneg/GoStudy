package main

import (
	"bufio"
	"os"
	"slices"

	"github.com/sergiovaneg/GoStudy/utils"
)

type Diagram [][]byte
type Coordinate [2]int

func (x Coordinate) isBounded(n, m int) bool {
	if x[0] < 0 || x[1] < 0 {
		return false
	}

	if x[0] >= n || x[1] >= m {
		return false
	}

	return true
}

func isLetter(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func (d Diagram) trace() (string, int) {
	x := Coordinate{0, slices.Index(d[0], '|')}
	dx := Coordinate{1, 0}
	n, m := len(d), len(d[0])
	var path string
	var steps int

	for x.isBounded(n, m) && d[x[0]][x[1]] != ' ' {
		steps++

		c := d[x[0]][x[1]]
		if c == '+' {
			var target byte

			if dx[0] == 0 {
				target = '|'
			} else {
				target = '-'
			}

			xp := Coordinate{x[0] + dx[1], x[1] - dx[0]}
			if xp.isBounded(n, m) && (d[xp[0]][xp[1]] == target || isLetter(d[xp[0]][xp[1]])) {
				dx = Coordinate{dx[1], -dx[0]}
			} else {
				dx = Coordinate{-dx[1], dx[0]}
			}
		} else if isLetter(c) {
			path += string(c)
		}

		x = Coordinate{x[0] + dx[0], x[1] + dx[1]}
	}

	return path, steps
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	d := make(Diagram, 0, n)
	for scanner.Scan() {
		d = append(d, []byte(scanner.Text()))
	}

	println(d.trace())
}
