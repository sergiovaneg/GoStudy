package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	knotHashUtils "github.com/sergiovaneg/GoStudy/AoC/2017/knotHash/knotHashUtils"
)

const n = 1 << 7

type Grid [n][n]bool

func bitRepresentation(hash string) [n]bool {
	var res [n]bool

	for i := range n >> 2 {
		mask, _ := strconv.ParseInt(hash[i:i+1], 16, 0)
		for j := range 4 {
			res[i<<2+j] = (mask>>(3-j))&0x1 == 0x1
		}
	}

	return res
}

func isValidCoordinate(i, j int) bool {
	if i < 0 || j < 0 {
		return false
	}
	if i >= n || j >= n {
		return false
	}

	return true
}

func (used *Grid) fill(i0, j0 int) {
	used[i0][j0] = false
	for _, dx := range [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		i1, j1 := i0+dx[0], j0+dx[1]
		if isValidCoordinate(i1, j1) && used[i1][j1] {
			used.fill(i1, j1)
		}
	}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	var used Grid
	for i := range n {
		keyStr := scanner.Text() + fmt.Sprintf("-%v", i)
		hash := knotHashUtils.KnotHash(keyStr)
		used[i] = bitRepresentation(hash)
	}

	var resA int
	for _, row := range used {
		for _, bit := range row {
			if bit {
				resA++
			}
		}
	}
	println(resA)

	var resB int
	for i := range n {
		for j := range n {
			if used[i][j] {
				resB++
				used.fill(i, j)
			}
		}
	}
	println(resB)
}
