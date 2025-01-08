package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

const W, H = 101, 103

type Robot struct {
	x0 [2]int
	v  [2]int
}
type robotSlice []Robot
type safetyArr [4]int

func moveNSteps(r Robot, n int) Robot {
	flags := [2]bool{r.v[0] < 0, r.v[1] < 0}

	x0, v := r.x0, r.v
	if flags[0] {
		x0[0] = W - 1 - x0[0]
		v[0] = -v[0]
	}
	if flags[1] {
		x0[1] = H - 1 - x0[1]
		v[1] = -v[1]
	}

	x := [2]int{
		(x0[0] + v[0]*n) % W,
		(x0[1] + v[1]*n) % H,
	}

	if flags[0] {
		x[0] = W - 1 - x[0]
	}
	if flags[1] {
		x[1] = H - 1 - x[1]
	}

	return Robot{
		x0: x,
		v:  r.v,
	}
}

func getQuadrantIdx(r Robot) int {
	if r.x0[0] < W>>1 && r.x0[1] < H>>1 {
		return 0
	}
	if r.x0[0] < W>>1 && r.x0[1] > H>>1 {
		return 1
	}
	if r.x0[0] > W>>1 && r.x0[1] < H>>1 {
		return 2
	}
	if r.x0[0] > W>>1 && r.x0[1] > H>>1 {
		return 3
	}

	return -1
}

func parseRobot(line string) Robot {
	var values [4]int
	for idx, num := range regexp.MustCompile(
		`-{0,1}\d+`).FindAllString(line, 4) {
		aux, _ := strconv.Atoi(num)
		values[idx] = aux
	}

	return Robot{
		x0: [2]int{values[0], values[1]},
		v:  [2]int{values[2], values[3]},
	}
}

func (sArr *safetyArr) updateSafety(r Robot) {
	if sIdx := getQuadrantIdx(moveNSteps(r, 100)); sIdx != -1 {
		sArr[sIdx]++
	}
}

func (sArr safetyArr) getSafety() (r int) {
	r = 1
	for _, s := range sArr {
		r *= s
	}
	return
}

func (rSlice robotSlice) printIfValid(step int) (found bool) {
	pat := make([][]byte, H)
	for i := range H {
		pat[i] = make([]byte, W)
		for j := range W {
			pat[i][j] = '.'
		}
	}
	for _, r := range rSlice {
		pat[r.x0[1]][r.x0[0]] = '#'
	}

	for _, line := range pat {
		if strings.Contains(string(line), strings.Repeat("#", 10)) {
			found = true
			break
		}
	}

	if found {
		file, _ := os.Create(fmt.Sprintf("./prints/%v.txt", step))
		defer file.Close()

		w := bufio.NewWriter(file)
		defer w.Flush()

		for _, line := range pat {
			fmt.Fprintln(w, string(line))
		}
	}

	return
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	var s safetyArr
	rArr := make(robotSlice, 0, n)
	for scanner.Scan() {
		r := parseRobot(scanner.Text())
		rArr = append(rArr, r)
		s.updateSafety(r)
	}

	println(s.getSafety())

	os.MkdirAll("./prints/", os.ModePerm)
	for step := 0; !rArr.printIfValid(step); step++ {
		for idx := range rArr {
			rArr[idx] = moveNSteps(rArr[idx], 1)
		}
	}
}
