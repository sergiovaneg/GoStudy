package main

import (
	"bufio"
	"log"
	"os"
	"slices"

	"github.com/sergiovaneg/GoStudy/utils"
)

const workerQuota = 128

type State [2][2]int
type Record map[State]bool

func (s State) canAdvance(lab [][]byte, obstacle [2]int) int {
	i, j := s[0][0]+s[1][0], s[0][1]+s[1][1]
	if i < 0 || j < 0 {
		return -1
	}

	m, n := len(lab), len(lab[0])
	if i == m || j == n {
		return -1
	}

	if lab[i][j] == '#' || [2]int{i, j} == obstacle {
		return 0
	}

	return 1
}

func (s State) takeStep(lab [][]byte, obstacle [2]int) State {
	switch s.canAdvance(lab, obstacle) {
	case 0:
		s[1][0], s[1][1] = s[1][1], -s[1][0]
	case 1:
		s[0][0], s[0][1] = s[0][0]+s[1][0], s[0][1]+s[1][1]
	default:
		s = State{}
	}
	return s
}

func generateRecord(s State, lab [][]byte, obstacle [2]int) Record {
	r := make(Record)
	for !r[s] {
		r[s] = true
		s = s.takeStep(lab, obstacle)
	}

	return r
}

func (r Record) getUnique() [][2]int {
	delete(r, State{})
	uniqueR := make(map[[2]int]bool, len(r))

	for s := range r {
		uniqueR[s[0]] = true
	}

	uniquePositions := make([][2]int, 0, len(uniqueR))
	for pos := range uniqueR {
		uniquePositions = append(uniquePositions, pos)
	}

	return uniquePositions
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
	r := generateRecord(s0, lab, [2]int{})
	uniquePos := r.getUnique()
	println(len(uniquePos))

	nCandidates := len(uniquePos)
	nPartitions := nCandidates / workerQuota
	if nCandidates%nPartitions > 0 {
		nPartitions++
	}

	c := make(chan int, nPartitions)
	for idx := 0; idx < nCandidates; idx += workerQuota {
		go func() {
			acc := 0
			for _, cand := range uniquePos[idx:min(idx+workerQuota, nCandidates)] {
				if rec := generateRecord(s0, lab, cand); !rec[State{}] {
					acc++
				}
			}
			c <- acc
		}()
	}

	res := 0
	for idx := 0; idx < nPartitions; idx++ {
		res += <-c
	}
	println(res)
}
