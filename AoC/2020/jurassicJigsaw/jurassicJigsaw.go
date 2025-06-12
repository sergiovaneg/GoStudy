package main

import (
	"os"
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

func (p Placement) getSide(ts TileSet, k int) []bool {
	tile := ts[p.id]
	side := make([]bool, len(tile))

	return side
}

func (a Arrangement) isValidCandidate(
	candidate Placement, ts TileSet, i, j int) bool

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
		println(splitBlock)

		tileSet[id] = parseTile(splitBlock[1:])
	}

	n := utils.ISqrt(len(tiles))
	arrangement := Arrangement{
		placedTiles: make(map[[2]int]Placement, n*n),
		record:      make(map[int]bool),
	}
	arrangement.recursiveSetup(tileSet, n, 0, 0)
}
