package main

import (
	"bufio"
	"os"
	"strings"
)

func passThrough(beams []int, row string) ([]int, int) {
	ret, cnt := make([]int, len(row)), 0

	for j, r := range row {
		nBeams := beams[j]
		if nBeams == 0 {
			continue
		}

		if r == '^' {
			ret[j-1] += nBeams
			ret[j+1] += nBeams
			cnt++
		} else {
			ret[j] += nBeams
		}
	}

	return ret, cnt
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	s0 := strings.Index(scanner.Text(), "S")
	beams := make([]int, len(scanner.Text()))
	beams[s0] = 1

	resA := 0
	for scanner.Scan() {
		var tmp int
		beams, tmp = passThrough(beams, scanner.Text())
		resA += tmp
	}
	println(resA)

	resB := 0
	for _, num := range beams {
		resB += num
	}
	println(resB)
}
