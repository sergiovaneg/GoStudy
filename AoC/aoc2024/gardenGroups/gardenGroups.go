package main

import (
	"bufio"
	"log"
	"os"

	"github.com/sergiovaneg/GoStudy/utils"
)

type Garden [][]byte
type IdMap [][]int

type WalkerState [2][2]int
type Walker struct {
	state  WalkerState
	record map[WalkerState]bool
	sides  int
}

func (im IdMap) isValidX(x [2]int) bool {
	i, j := x[0], x[1]
	if i < 0 || j < 0 {
		return false
	}
	if i >= len(im) || j >= len(im[i]) {
		return false
	}
	return true
}

func (im IdMap) isUnmarked(x [2]int) bool {
	return im.isValidX(x) && im[x[0]][x[1]] == 0
}

func (im IdMap) recursiveMark(g Garden, x0 [2]int, id int) {
	im[x0[0]][x0[1]] = id
	for _, dx := range [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
		x := [2]int{x0[0] + dx[0], x0[1] + dx[1]}
		if im.isUnmarked(x) && g[x[0]][x[1]] == g[x0[0]][x0[1]] {
			im.recursiveMark(g, x, id)
		}
	}
}

func (im IdMap) isInsideGroup(x [2]int, id int) bool {
	return im.isValidX(x) && im[x[0]][x[1]] == id
}

func (im IdMap) getAreaPerimeters() map[int][2]int {
	r := make(map[int][2]int)
	for i, row := range im {
		for j, id := range row {
			ap := r[id]
			ap[0]++
			for _, dx := range [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {-1, 0}} {
				x := [2]int{i + dx[0], j + dx[1]}
				if !im.isInsideGroup(x, id) {
					ap[1]++
				}
			}
			r[id] = ap
		}
	}
	return r
}

func spawnWalker(i, j int) Walker {
	return Walker{
		state:  WalkerState{{i, j}, {-1, 0}},
		record: make(map[WalkerState]bool),
		sides:  0,
	}
}

func (w Walker) canMove(im IdMap, dir [2]int) bool {
	x := [2]int{w.state[0][0] + dir[0], w.state[0][1] + dir[1]}
	return im.isInsideGroup(x, im[w.state[0][0]][w.state[0][1]])
}

func (w *Walker) takeStep(im IdMap) bool {
	if w.record[w.state] {
		return false
	}
	w.record[w.state] = true

	if dL := [2]int{
		-w.state[1][1],
		w.state[1][0],
	}; w.canMove(im, dL) {
		w.sides++
		w.state[1] = dL
		w.state = WalkerState{
			{
				w.state[0][0] + dL[0],
				w.state[0][1] + dL[1],
			},
			dL,
		}
	} else if w.canMove(im, w.state[1]) {
		w.state[0] = [2]int{
			w.state[0][0] + w.state[1][0],
			w.state[0][1] + w.state[1][1],
		}
	} else {
		w.sides++
		w.state[1] = [2]int{
			w.state[1][1],
			-w.state[1][0],
		}
	}

	return true
}

func (im IdMap) getSides() map[int]int {
	charted := make(map[int]bool)
	c := make(chan [2]int)

	for i, row := range im {
		for j, id := range row {
			if charted[id] {
				continue
			}
			charted[id] = true

			go func(w Walker, id int) {
				// Walker starts going up
				for w.canMove(im, w.state[1]) {
					w.state[0][0]--
				}

				// Go right for left-hand rule
				w.state[1] = [2]int{0, 1}

				for w.takeStep(im) {
				}
				c <- [2]int{id, w.sides}
			}(spawnWalker(i, j), id)
		}
	}

	r := make(map[int]int)
	for range charted {
		aux := <-c
		r[aux[0]] = aux[1]
	}

	return r
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	im := make(IdMap, 0, n)
	g := make(Garden, 0, n)
	for scanner.Scan() {
		aux := scanner.Text()
		im = append(im, make([]int, len(aux)))
		g = append(g, []byte(aux))
	}

	id := 0
	for i, row := range g {
		for j := range row {
			x0 := [2]int{i, j}
			if !im.isUnmarked(x0) {
				continue
			}
			id++
			im.recursiveMark(g, x0, id)
		}
	}

	resA, aps := 0, im.getAreaPerimeters()
	for _, ap := range aps {
		resA += ap[0] * ap[1]
	}
	println(resA)

	resB, sides := 0, im.getSides()
	for id := range sides {
		resB += sides[id] * aps[id][0]
	}
	println(resB)
}
