package main

import (
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

type Tile [][]bool
type TileSet map[int]Tile

type Placement struct {
	id    int
	state int
}
type Arrangement struct {
	placedTiles map[[2]int]Placement
	record      map[int]bool
}

func parseTile(lines []string) Tile {
	n := len(lines)
	t := make(Tile, n)

	for i, line := range lines {
		row := make([]bool, n)
		for j, val := range line {
			row[j] = val == '#'
		}
		t[i] = row
	}

	return t
}

func (t Tile) getSide(k int) []bool {
	n := len(t)
	side := make([]bool, n)
	switch k {
	case 0:
		for j := range n {
			side[j] = t[0][j]
		}
	case 1:
		for i := range n {
			side[i] = t[i][n-1]
		}
	case 2:
		for j := range n {
			side[j] = t[n-1][n-j-1]
		}
	case 3:
		for i := range n {
			side[i] = t[n-i-1][0]
		}
	}
	return side
}

func (p Placement) getSide(ts TileSet, k int) []bool {
	var side []bool
	tile := ts[p.id]
	normIdx := (p.state + k) & 0b011

	if p.state&0b100 == 0 {
		side = tile.getSide(normIdx)
	} else {
		side = tile.getSide(0b11 - normIdx)
		slices.Reverse(side)
	}

	return side
}

func (a Arrangement) isValidCandidate(
	candidate Placement, ts TileSet, i, j int) bool {
	if i > 0 {
		topSide := a.placedTiles[[2]int{i - 1, j}].getSide(ts, 2)
		slices.Reverse(topSide)
		if !slices.Equal(candidate.getSide(ts, 0), topSide) {
			return false
		}
	}

	if j > 0 {
		leftSide := a.placedTiles[[2]int{i, j - 1}].getSide(ts, 1)
		slices.Reverse(leftSide)
		if !slices.Equal(candidate.getSide(ts, 3), leftSide) {
			return false
		}
	}

	return true
}

func (a *Arrangement) recursiveSetup(ts TileSet, n int, i, j int) bool {
	for tsId := range ts {
		if a.record[tsId] {
			continue
		}

		a.record[tsId] = true
		for state := range 8 {
			candidate := Placement{id: tsId, state: state}
			if !a.isValidCandidate(candidate, ts, i, j) {
				continue
			}
			a.placedTiles[[2]int{i, j}] = candidate

			if (i == j && i == n-1) || a.recursiveSetup(
				ts, n, i+(j+1)/n, (j+1)%n) {
				return true
			}
		}
		a.record[tsId] = false
	}
	return false
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	tiles := strings.Split(strings.Trim(string(data), "\n"), "\n\n")

	tileSet := make(TileSet)
	for _, block := range tiles {
		splitBlock := strings.Split(block, "\n")

		id, _ := strconv.Atoi(splitBlock[0][5 : len(splitBlock[0])-1])

		tileSet[id] = parseTile(splitBlock[1:])
	}

	n := utils.ISqrt(len(tiles))
	arrangement := Arrangement{
		placedTiles: make(map[[2]int]Placement, n*n),
		record:      make(map[int]bool),
	}
	valid := arrangement.recursiveSetup(tileSet, n, 0, 0)

	if valid {
		resA := 1
		for _, i := range [2]int{0, n - 1} {
			for _, j := range [2]int{0, n - 1} {
				resA *= arrangement.placedTiles[[2]int{i, j}].id
			}
		}
		println(resA)
	} else {
		panic("Error")
	}
}
