package main

import (
	"bufio"
	"log"
	"os"

	"github.com/sergiovaneg/GoStudy/utils"
)

type WS [][]byte

func (ws *WS) safeGet(i, j int) byte {
	if i < 0 || j < 0 {
		return '.'
	}

	if i >= len(*ws) {
		return '.'
	}

	if j >= len((*ws)[i]) {
		return '.'
	}

	return (*ws)[i][j]
}

func (ws *WS) checkStraight(i, j int) int {
	cnt := 0
	for _, d := range [8][2]int{
		{-1, -1}, // NW
		{-1, 0},  // N
		{-1, 1},  // NE
		{0, -1},  // W
		{0, 1},   // E
		{1, -1},  // SW
		{1, 0},   // S
		{1, 1},   // SE
	} {
		m := ws.safeGet(i+d[0], j+d[1])
		a := ws.safeGet(i+2*d[0], j+2*d[1])
		s := ws.safeGet(i+3*d[0], j+3*d[1])
		if m == 'M' && a == 'A' && s == 'S' {
			cnt++
		}
	}
	return cnt
}

func (ws *WS) checkCross(i, j int) bool {
	if a, b := ws.safeGet(i-1, j-1), ws.safeGet(i+1, j+1); a+b == 'M'+'S' {
		if a, b = ws.safeGet(i-1, j+1), ws.safeGet(i+1, j-1); a+b == 'M'+'S' {
			return true
		}
		return false
	}
	return false
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	lines := make(WS, 0, n)
	for scanner.Scan() {
		lines = append(lines, []byte(scanner.Text()))
	}

	res_0 := 0
	for i, line := range lines {
		for j, g := range line {
			if g == 'X' {
				res_0 += lines.checkStraight(i, j)
			}
		}
	}
	println(res_0)

	res_1 := 0
	for i, line := range lines {
		for j, g := range line {
			if g == 'A' && lines.checkCross(i, j) {
				res_1++
			}
		}
	}
	println(res_1)
}
