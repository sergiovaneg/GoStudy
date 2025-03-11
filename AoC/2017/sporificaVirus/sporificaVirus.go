package main

import (
	"bufio"
	"os"

	"github.com/sergiovaneg/GoStudy/utils"
)

const BurstsNaive = 10000
const BurstsEvolved = 10000000

type Coordinate [2]int
type Network map[Coordinate]int
type Carrier struct {
	x   Coordinate
	dx  Coordinate
	net Network
	cnt int
}

func newCarrier(rows []string) Carrier {
	n, m := len(rows), len(rows[0])
	net := make(Network)
	for i, row := range rows {
		for j, r := range row {
			if r == '#' {
				net[Coordinate{i, j}] = 2
			}
		}
	}

	return Carrier{
		x:   [2]int{n >> 1, m >> 1},
		dx:  [2]int{-1, 0},
		net: net,
		cnt: 0,
	}
}

func (c *Carrier) iterateNaive() {
	if c.net[c.x] == 2 {
		c.dx = Coordinate{c.dx[1], -c.dx[0]}
		c.net[c.x] = 0
	} else {
		c.dx = Coordinate{-c.dx[1], c.dx[0]}
		c.net[c.x] = 2
		c.cnt++
	}

	c.x[0] += c.dx[0]
	c.x[1] += c.dx[1]
}

func (c *Carrier) iterateEvolved() {
	switch c.net[c.x] {
	case 0:
		c.dx = Coordinate{-c.dx[1], c.dx[0]}
	case 1:
		c.cnt++
	case 2:
		c.dx = Coordinate{c.dx[1], -c.dx[0]}
	case 3:
		c.dx = Coordinate{-c.dx[0], -c.dx[1]}
	}

	c.net[c.x] = (c.net[c.x] + 1) % 4

	c.x[0] += c.dx[0]
	c.x[1] += c.dx[1]
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	rows := make([]string, 0, n)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	carrier := newCarrier(rows)
	for range BurstsNaive {
		carrier.iterateNaive()
	}
	println(carrier.cnt)

	carrier = newCarrier(rows)
	for range BurstsEvolved {
		carrier.iterateEvolved()
	}
	println(carrier.cnt)
}
