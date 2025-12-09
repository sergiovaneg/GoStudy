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

func (p path) filterEndpoints(idx int) []int {
	n := len(p)

	if idx >= n-3 {
		return []int{}
	}

	var x, y, z coord
	y = p[idx]

	if idx == 0 {
		x, z = p[n-1], p[1]
	} else {
		x, z = p[idx-1], p[idx+1]
	}

	d0, d1 := delta(x, y), delta(z, y)

	var f func(coord) bool
	if cross(d0, d1) > 0 { // left turn (90° inclusive)
		f = func(c coord) bool {
			dc := delta(c, y)
			dot0, dot1 := dot(dc, d0), dot(dc, d1)
			return dot0 > 0 && dot1 > 0
		}
	} else { //right turn (90° exclusive)
		f = func(c coord) bool {
			dc := delta(c, y)
			dot0, dot1 := dot(dc, d0), dot(dc, d1)
			return dot0 < 0 || dot1 < 0
		}
	}

	ret := make([]int, 0)
	for offset, c := range p[idx+2:] {
		if f(c) {
			ret = append(ret, idx+2+offset)
		}
	}

	return ret
}

func (p path) getValidArea(i, j int) int {
	bounds := [2][2]int{
		{min(p[i][0], p[j][0]), max(p[i][0], p[j][0])},
		{min(p[i][1], p[j][1]), max(p[i][1], p[j][1])},
	}

	var cnstIdx, varIdx int
	for k := 0; k < len(p)-1; k++ {
		v, w := p[k], p[k+1]

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

	p := make(path, 0, n+1)

	for scanner.Scan() {
		p = append(p, parseCoord(scanner.Text()))
	}
	p = append(p, p[0])

	if !p.isClockwise() {
		slices.Reverse(p)
	}

	var resA, resB int
	for i := range len(p) - 1 {
		for offset := range p[i+1:] {
			resA = max(resA, p.getRectangleArea(i, i+1+offset))
		}

		for _, j := range p.filterEndpoints(i) {
			if p.getRectangleArea(i, j) <= resB {
				continue
			}
			resB = max(resB, p.getValidArea(i, j))
		}
	}

	println(resA)
	println(resB)
}
