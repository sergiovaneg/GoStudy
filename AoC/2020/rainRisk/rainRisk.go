package main

import (
	"bufio"
	"os"
	"strconv"

	"github.com/sergiovaneg/GoStudy/utils"
)

type coordinate [2]int
type ship struct {
	x, dx       coordinate
	cardinalPtr *coordinate
}

func initShip(di, dj int, isRelative bool) *ship {
	s := ship{dx: coordinate{di, dj}}
	if isRelative {
		s.cardinalPtr = &s.dx
	} else {
		s.cardinalPtr = &s.x
	}
	return &s
}

func (s *ship) update(inst string) {
	val, _ := strconv.Atoi(inst[1:])

	switch inst[0] {
	case 'N':
		s.cardinalPtr[0] -= val
	case 'S':
		s.cardinalPtr[0] += val
	case 'E':
		s.cardinalPtr[1] += val
	case 'W':
		s.cardinalPtr[1] -= val
	case 'F':
		s.x[0] += val * s.dx[0]
		s.x[1] += val * s.dx[1]
	case 'L':
		for range (val / 90) % 4 {
			s.dx = coordinate{-s.dx[1], s.dx[0]}
		}
	case 'R':
		for range (val / 90) % 4 {
			s.dx = coordinate{s.dx[1], -s.dx[0]}
		}
	}
}

func manhattan(x coordinate) int {
	return utils.AbsInt(x[0]) + utils.AbsInt(x[1])
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	instSet := make([]string, 0, n)

	for scanner.Scan() {
		instSet = append(instSet, scanner.Text())
	}

	sA, sB := initShip(0, 1, false), initShip(-1, 10, true)
	for _, inst := range instSet {
		sA.update(inst)
		sB.update(inst)
	}
	println(manhattan(sA.x))
	println(manhattan(sB.x))
}
