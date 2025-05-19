package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

type Tile struct {
	rot   int
	sides [4][]bool
}

func (t Tile) getSide(i int) []bool {
	return t.sides[(i+t.rot)%4]
}

func parseTile(data string) []Tile {
	refTile := Tile{}
	for line := range strings.SplitSeq(data, "\n") {

	}
}

func recursiveAssemble(
	tileSets map[int][]Tile,
	idx, n int,
	used *map[int]bool,
	conf *map[[2]int][2]int, // Id and equivalenceIdx
) bool {
	if idx == 0 {
		*used = make(map[int]bool)
		*conf = make(map[[2]int][2]int)
	} else if idx == n*n {
		return true
	}
	x := [2]int{idx / n, idx % n}
	for tileId, tileSet := range tileSets {
		if (*used)[tileId] {
			continue
		}

		var targetLeft, targetTop []bool
		if idRot, ok := (*conf)[[2]int{x[0], x[1] - 1}]; ok {
			targetLeft = tileSets[idRot[0]][idRot[1]].getSide(1)
		}
		if idRot, ok := (*conf)[[2]int{x[0] - 1, x[1]}]; ok {
			targetTop = tileSets[idRot[0]][idRot[1]].getSide(2)
		}

		(*used)[tileId] = true
		for setIdx, tile := range tileSet {
			if targetLeft != nil && !slices.Equal(targetLeft, tile.getSide(3)) {
				continue
			}

			if targetTop != nil && !slices.Equal(targetLeft, tile.getSide(3)) {
				continue
			}

			(*conf)[x] = [2]int{tileId, setIdx}
			if recursiveAssemble(tileSets, idx+1, n, used, conf) {
				return true
			}
		}
		(*used)[tileId] = false
	}

	return false
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	tiles := strings.Split(string(data), "\nTile ")

	n := utils.ISqrt(len(tiles))

	tileSets := make(map[int][]Tile)
	for _, block := range tiles {
		splitBlock := strings.Split(block, "\n")

		l := len(splitBlock[0])
		id, _ := strconv.Atoi(splitBlock[0][:l-1])

		tileSets[id] = parseTile(splitBlock[1])
	}

	conf := new(map[[2]int][2]int)
	if recursiveAssemble(tileSets, 0, n, new(map[int]bool), conf) {
		for x, idRot := range *conf {
			fmt.Printf("(%v,%v): %v (%v)\n", x[0], x[1], idRot[0], idRot[1])
		}
	}
}
