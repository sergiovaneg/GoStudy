package main

import (
	"bufio"
	"log"
	"os"
	"sync"

	"github.com/sergiovaneg/GoStudy/utils"
)

const rotCost = 1000
const stepCost = 1

type Coordinate [2]int
type State [2]Coordinate
type Record map[State]int
type Maze [][]byte

type ParallelRecord struct {
	mu  sync.Mutex
	rec Record
	m   Maze
}

func (pr *ParallelRecord) getCandidates(s State) []State {
	candidates := make([]State, 0)
	x0 := s[0]
	for _, dx := range [4]Coordinate{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	} {
		x1 := Coordinate{x0[0] + dx[0], x0[1] + dx[1]}
		if pr.m[x1[0]][x1[1]] != '#' {
			candidates = append(candidates, State{x1, dx})
		}
	}

	return candidates
}

func getTransitionCost(s0, s1 State) int {
	v0, v1 := s0[1], s1[1]
	switch v0[0]*v1[0] + v0[1]*v1[1] {
	case 1:
		return stepCost
	case 0:
		return rotCost + stepCost
	default:
		return rotCost<<1 + stepCost
	}
}

func (pr *ParallelRecord) parallelDFS(s0 State, accCost int) {
	var wg sync.WaitGroup

	for {
		pr.mu.Lock()
		currentBest, exists := pr.rec[s0]
		if !exists || accCost < currentBest {
			pr.rec[s0] = accCost
			pr.mu.Unlock()
		} else {
			pr.mu.Unlock()
			break
		}

		candidates := pr.getCandidates(s0)
		for _, s1 := range candidates[1:] {
			wg.Add(1)
			go func(newS State, newCost int) {
				pr.parallelDFS(newS, newCost)
				wg.Done()
			}(s1, accCost+getTransitionCost(s0, s1))
		}

		accCost += getTransitionCost(s0, candidates[0])
		s0 = candidates[0]
	}

	wg.Wait()
}

func (m Maze) getLandmarks() (start, end Coordinate) {
	for i, row := range m {
		for j, val := range row {
			if val == 'S' {
				start = Coordinate{i, j}
			} else if val == 'E' {
				end = Coordinate{i, j}
			}
		}
	}
	return
}

func (rec Record) getMinScore(end Coordinate) int {
	minVal := -1

	for k, v := range rec {
		if k[0] != end {
			continue
		}
		if minVal == -1 || v < minVal {
			minVal = v
		}
	}

	return minVal
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	var pr ParallelRecord
	pr.rec = make(Record)
	pr.m = make(Maze, n)

	for i := 0; scanner.Scan(); i++ {
		pr.m[i] = []byte(scanner.Text())
	}

	s, e := pr.m.getLandmarks()
	pr.parallelDFS(State{s, Coordinate{0, -1}}, 0)
	println(pr.rec.getMinScore(e))
}
