package main

import (
	"bufio"
	"log"
	"os"
	"slices"

	"github.com/sergiovaneg/GoStudy/utils"
)

const pathLen = 10

type Coord [2]int
type Path [pathLen]Coord
type TMap [][]rune

func (m TMap) getValidNeighbours(c Coord) []Coord {
	res := make([]Coord, 0)
	for _, d := range [4]Coord{
		{0, -1}, {0, 1}, {-1, 0}, {1, 0},
	} {
		i, j := c[0]+d[0], c[1]+d[1]

		if i < 0 || j < 0 {
			continue
		}
		if i >= len(m) || j >= len(m[0]) {
			continue
		}

		if m[i][j] == m[c[0]][c[1]]+1 {
			res = append(res, Coord{i, j})
		}
	}
	return res
}

func (m TMap) traversePath(n int, p Path) []Path {
	if n == pathLen-1 {
		return []Path{p}
	}

	res := make([]Path, 0)
	for _, x1 := range m.getValidNeighbours(p[n]) {
		p[n+1] = x1
		res = append(res, m.traversePath(n+1, p)...)
	}
	return res
}

func getUniqueTails(paths []Path) []Coord {
	res := make([]Coord, 0, len(paths))
	for _, p := range paths {
		if !slices.Contains(res, p[pathLen-1]) {
			res = append(res, p[pathLen-1])
		}
	}
	return res
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	x0Arr := make([]Coord, 0)
	m := make(TMap, n)
	for i := 0; scanner.Scan(); i++ {
		m[i] = []rune(scanner.Text())
		for j, r := range m[i] {
			if r == '0' {
				x0Arr = append(x0Arr, Coord{i, j})
			}
		}
	}

	c := make(chan []Path, len(x0Arr))
	totalScore, totalRating := 0, 0
	for _, x0 := range x0Arr {
		go func() {
			c <- m.traversePath(0, Path{x0})
		}()
	}
	for range x0Arr {
		paths := <-c
		totalScore += len(getUniqueTails(paths))
		totalRating += len(paths)
	}
	close(c)

	println(totalScore)
	println(totalRating)
}
