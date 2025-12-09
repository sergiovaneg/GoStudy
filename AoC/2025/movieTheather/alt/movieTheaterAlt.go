package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

type coord [2]int
type path []coord

func parseCoord(line string) coord {
	var x coord
	nums := strings.Split(line, ",")
	x[0], _ = strconv.Atoi(nums[0])
	x[1], _ = strconv.Atoi(nums[1])
	return x
}

func (p path) getRectangleArea(i, j int) int {
	x, y := p[i], p[j]
	return (utils.AbsInt(x[0]-y[0]) + 1) * (utils.AbsInt(x[1]-y[1]) + 1)
}

func delta(v, w coord) coord {
	return coord{v[0] - w[0], v[1] - w[1]}
}

func cross(v, w coord) int {
	return v[0]*w[1] - v[1]*w[0]
}

func dot(v, w coord) int {
	return v[0]*w[0] + v[1]*w[1]
}

func (p path) isClockwise() bool {
	cnt := 0
	n := len(p)
	for i := range n - 2 {
		d0 := delta(p[i], p[i+1])
		d1 := delta(p[i+2], p[i+1])
		if cross(d0, d1) > 0 {
			cnt++
		} else {
			cnt--
		}
	}

	return cnt > 0
}

func (p path) filterEndpoints(i int) []int {
	n := len(p)

	x, y, z := p[i-1], p[i], p[i+1]
	d0, d1 := delta(x, y), delta(z, y)

	var f func(coord) bool
	if cross(d0, d1) > 0 { // left turn (90° inclusive)
		f = func(c coord) bool {
			dc := delta(c, y)
			dot0, dot1 := dot(dc, d0), dot(dc, d1)
			return dot0 >= 0 && dot1 >= 0
		}
	} else { //right turn (90° exclusive)
		f = func(c coord) bool {
			dc := delta(c, y)
			dot0, dot1 := dot(dc, d0), dot(dc, d1)
			return dot0 <= 0 || dot1 <= 0
		}
	}

	ret := make([]int, 0)
	for j, c := range p[1 : n-1] {
		if i == j+1 {
			continue
		}
		if f(c) {
			ret = append(ret, j+1)
		}
	}

	return ret
}

func (p path) getValidArea(i, j int, pFilt []int) int {
	bounds := [2][2]int{
		{min(p[i][0], p[j][0]), max(p[i][0], p[j][0])},
		{min(p[i][1], p[j][1]), max(p[i][1], p[j][1])},
	}

	for _, k := range pFilt {
		v := p[k]
		for _, dk := range []int{-1, 1} {
			w := p[k+dk]

			var cnstIdx, varIdx int
			if v[0] == w[0] { // horizontal wall
				cnstIdx, varIdx = 0, 1
			} else { // vertical wall
				cnstIdx, varIdx = 1, 0
			}

			if v[cnstIdx] <= bounds[cnstIdx][0] || v[cnstIdx] >= bounds[cnstIdx][1] {
				continue
			}
			if max(v[varIdx], w[varIdx]) <= bounds[varIdx][0] {
				continue
			}
			if min(v[varIdx], w[varIdx]) >= bounds[varIdx][1] {
				continue
			}
			return 0
		}
	}

	return p.getRectangleArea(i, j)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	p := make(path, 0, n+2)

	for scanner.Scan() {
		p = append(p, parseCoord(scanner.Text()))
	}

	first, last := p[0], p[n-1]
	p = append(p, first)
	p = slices.Insert(p, 0, last)

	if !p.isClockwise() {
		slices.Reverse(p)
	}

	var resA, resB int
	for i := range len(p) - 2 {
		for offset := range len(p) - 3 - i {
			resA = max(resA, p.getRectangleArea(i+1, i+2+offset))
		}

		pFilt := p.filterEndpoints(i + 1)
		for _, j := range pFilt {
			if i+1 >= j {
				continue
			}
			if p.getRectangleArea(i+1, j) <= resB {
				continue
			}
			resB = max(resB, p.getValidArea(i, j, pFilt))
		}
	}

	println(resA)
	println(resB)
}
