package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const memsize int = 71
const bytelim int = 1024

type Coord [2]int
type Corrupted map[Coord]bool
type Djikstra map[Coord]int

func parseCoord(line string) Coord {
	var c Coord
	for i, num := range strings.Split(line, ",") {
		c[i], _ = strconv.Atoi(num)
	}
	return c
}

func (corr Corrupted) getNeighbours(x Coord) []Coord {
	neighbours := make([]Coord, 0)

	for _, dx := range [4]Coord{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	} {
		cand := Coord{x[0] + dx[0], x[1] + dx[1]}

		if cand[0] < 0 || cand[0] >= memsize {
			continue
		}
		if cand[1] < 0 || cand[1] >= memsize {
			continue
		}
		if corr[cand] {
			continue
		}

		neighbours = append(neighbours, cand)
	}

	return neighbours
}

func (corr Corrupted) shortestPath(start, end Coord) int {
	d := make(Djikstra)
	d[start] = 0

	queue, dist := []Coord{start}, 0

	for len(queue) > 0 {
		dist++
		newQueue := make([]Coord, 0)

		for _, src := range queue {
			for _, dst := range corr.getNeighbours(src) {
				if _, ok := d[dst]; ok {
					continue
				}

				newQueue = append(newQueue, dst)
				d[dst] = dist
			}
		}

		if _, ok := d[end]; ok {
			break
		}

		queue = newQueue
	}

	if steps, ok := d[end]; ok {
		return steps
	} else {
		return -1
	}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	corr := make(Corrupted)
	for n := 0; scanner.Scan() && n < bytelim; n++ {
		corr[parseCoord(scanner.Text())] = true
	}

	s, e := Coord{0, 0}, Coord{memsize - 1, memsize - 1}
	println(corr.shortestPath(s, e))

	for ok := true; ok; ok = corr.shortestPath(s, e) != -1 && scanner.Scan() {
		corr[parseCoord(scanner.Text())] = true
	}
	println(scanner.Text())
}
