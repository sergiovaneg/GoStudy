package main

import (
	"bufio"
	"log"
	"os"
	"slices"

	"github.com/sergiovaneg/GoStudy/utils"
)

type Coordinate [2]int
type Warehouse map[Coordinate]bool

type State struct {
	x Coordinate
	w Warehouse
	b Warehouse
}

func parseCharacter(c rune) Coordinate {
	switch c {
	case '^':
		return Coordinate{-1, 0}
	case '<':
		return Coordinate{0, -1}
	case '>':
		return Coordinate{0, 1}
	case 'v':
		return Coordinate{1, 0}
	default:
		return Coordinate{}
	}
}

func (x Coordinate) offset(dx Coordinate) Coordinate {
	return Coordinate{x[0] + dx[0], x[1] + dx[1]}
}

func (wh Warehouse) findCollisions(xs []Coordinate, w0, w1 int) []Coordinate {
	coll := make([]Coordinate, 0)
	for _, x0 := range xs {
		for n := 0; n > -w1; n-- { // Range over obstacle width
			if x := x0.offset(Coordinate{0, n}); wh[x] {
				coll = append(coll, x)
				break
			}
		}
		for n := 1; n < w0; n++ { // Range over self width
			if x := x0.offset(Coordinate{0, n}); wh[x] {
				coll = append(coll, x)
				break
			}
		}
	}

	return coll
}

func (s *State) execute(c rune, mapWidth int) {
	xs, dx := []Coordinate{s.x}, parseCharacter(c)
	impacted := make([]Coordinate, 0)

	width := 1
	for {
		for i := range xs {
			xs[i] = xs[i].offset(dx)
		}

		if wColl := s.w.findCollisions(xs, width, mapWidth); len(wColl) != 0 {
			impacted = nil
			break
		} else if bColl := utils.SliceDifference(
			s.b.findCollisions(
				xs, width, mapWidth),
			impacted); len(bColl) != 0 {
			// If there was a box collision, use the map width as self width
			width = mapWidth
			impacted = append(impacted, bColl...)

			xs = bColl
		} else {
			break
		}
	}

	if impacted != nil {
		s.x = s.x.offset(dx)
		slices.Reverse(impacted)
		for _, x := range impacted {
			s.b[x] = false
			x = x.offset(dx)
			s.b[x] = true
		}
	}
}

func (s State) getGpsSum() (r int) {
	r = 0
	for k, v := range s.b {
		if !v {
			continue
		}
		r += 100*k[0] + k[1]
	}
	return
}

/* func (s State) print(width int) {
	//Init
	var h, w int
	for k, v := range s.w {
		if !v {
			continue
		}

		if k[0] > h {
			h = k[0]
		}
		if k[1] > w {
			w = k[1]
		}
	}
	h++
	w += width

	// Grid preparation
	grid := make([][]byte, h)
	for i := range h {
		grid[i] = make([]byte, w)
		for j := range w {
			grid[i][j] = '.'
		}
	}

	// Wall inscription
	for wall := range s.w {
		for range width {
			grid[wall[0]][wall[1]] = '#'
			wall = wall.offset(Coordinate{0, 1})
		}
	}

	// Inner-Loop creation
	var innerLoop func(Coordinate)
	if width == 1 {
		innerLoop = func(x Coordinate) {
			grid[x[0]][x[1]] = 'O'
		}
	} else {
		innerLoop = func(x Coordinate) {
			grid[x[0]][x[1]] = '['
			for k := 1; k < width-1; k++ {
				x = x.offset(Coordinate{0, 1})
				grid[x[0]][x[1]] = '='
			}
			x = x.offset(Coordinate{0, 1})
			grid[x[0]][x[1]] = ']'
		}
	}

	// Box inscription
	for box, current := range s.b {
		if current {
			innerLoop(box)
		}
	}

	// Console print
	for _, row := range grid {
		println(string(row))
	}
	println()
} */

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var s, sWide State
	s.b, sWide.b = make(Warehouse), make(Warehouse)
	s.w, sWide.w = make(Warehouse), make(Warehouse)

	for i := 0; scanner.Scan(); i++ {
		if scanner.Text() == "" {
			break
		}

		for j, c := range scanner.Text() {
			if c == '#' {
				x := Coordinate{i, j}
				s.w[x] = true
				x[1] <<= 1
				sWide.w[x] = true
			} else if c == 'O' {
				s.b[Coordinate{i, j}] = true
				sWide.b[Coordinate{i, j << 1}] = true
			} else if c == '@' {
				s.x = Coordinate{i, j}
				sWide.x = Coordinate{i, j << 1}
			}
		}
	}

	for scanner.Scan() {
		for _, c := range scanner.Text() {
			s.execute(c, 1)
			sWide.execute(c, 2)
		}
	}

	println(s.getGpsSum())
	println(sWide.getGpsSum())
}
