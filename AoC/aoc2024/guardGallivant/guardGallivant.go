package main

import (
	"bufio"
	"log"
	"os"
	"slices"

	"github.com/sergiovaneg/GoStudy/utils"
)

type State [2][2]int
type Record map[State]int

func (s State) canAdvance(lab [][]byte) int {
	i, j := s[0][0]+s[1][0], s[0][1]+s[1][1]
	if i < 0 || j < 0 {
		return -1
	}

	m, n := len(lab), len(lab[0])
	if i == m || j == n {
		return -1
	}

	if lab[i][j] == '#' {
		return 0
	}

	return 1
}

func (s State) takeStep(lab [][]byte) State {
	flag := s.canAdvance(lab)

	if flag == -1 {
		return State{}
	}

	if flag == 1 {
		return State{
			[2]int{s[0][0] + s[1][0], s[0][1] + s[1][1]},
			s[1],
		}
	}

	s[1][0], s[1][1] = s[1][1], -s[1][0]
	return s
}

func generateRecord(s State, lab [][]byte) Record {
	r := make(Record)
	for r[s] < 1 {
		r[s]++
		s = s.takeStep(lab)
	}

	return r
}

func (r Record) countUnique() int {
	delete(r, State{})
	uniquePositions := make(map[[2]int]bool, len(r))

	for s := range r {
		uniquePositions[s[0]] = true
	}

	return len(uniquePositions)
}

func parseInitialState(lab [][]byte) State {
	for i, row := range lab {
		if j := slices.Index(row, '^'); j != -1 {
			return State{[2]int{i, j}, {-1, 0}}
		}
	}

	return State{}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	lab := make([][]byte, 0, n)
	for scanner.Scan() {
		lab = append(lab, []byte(scanner.Text()))
	}

	s0 := parseInitialState(lab)
	r := generateRecord(s0, lab)
	println(r.countUnique())

	c := make(chan int, len(lab))
	for i := range lab {
		labMirror := make([][]byte, len(lab))
		for j, row := range lab {
			labMirror[j] = make([]byte, len(row))
			copy(labMirror[j], row)
		}
		go func() {
			res := 0
			for j := range labMirror[i] {
				if labMirror[i][j] == '.' {
					labMirror[i][j] = '#'
					if r = generateRecord(s0, labMirror); r[State{}] == 0 {
						res++
					}
					labMirror[i][j] = '.'
				}
			}
			c <- res
		}()
	}

	res_1 := 0
	for range lab {
		res_1 += <-c
	}
	close(c)
	println(res_1)
}
