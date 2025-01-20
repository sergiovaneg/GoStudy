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
type Maze [][]byte

type Record map[State]int
type Traversed map[Coordinate]bool

type ParallelRecord struct {
	best int
	mu   *sync.Mutex
	rec  Record
	m    Maze
	t    Traversed
}

type Walker struct {
	acc int
	s   State
	t   Traversed
}

func initWalker(acc int, s State, t Traversed) Walker {
	newT := make(Traversed)
	for k := range t {
		newT[k] = true
	}
	newT[s[0]] = true

	return Walker{acc: acc, s: s, t: newT}
}

func (tSelf *Traversed) merge(tOther Traversed) {
	for k := range tOther {
		(*tSelf)[k] = true
	}
}

func (pr ParallelRecord) getCandidateStates(s State) []State {
	candidates := make([]State, 0)
	x0 := s[0]
	for _, dx := range [4]Coordinate{
		{0, -1}, {0, 1}, {-1, 0}, {1, 0},
	} {
		x1 := Coordinate{x0[0] + dx[0], x0[1] + dx[1]}
		if pr.m[x1[0]][x1[1]] == '#' {
			continue
		}
		candidates = append(candidates, State{x1, dx})
	}

	return candidates
}

func getTransitionCost(s0, s1 State) int {
	v0, v1 := s0[1], s1[1]
	if dot := v0[0]*v1[0] + v0[1]*v1[1]; dot > 0 {
		return stepCost
	} else if dot < 0 {
		return rotCost<<1 + stepCost
	} else {
		return rotCost + stepCost
	}
}

func (pr *ParallelRecord) parallelBFS(w Walker, end Coordinate) {
	var wg sync.WaitGroup

	for {
		// Atomic part
		pr.mu.Lock()

		if pr.best != -1 && w.acc > pr.best {
			pr.mu.Unlock()
			break
		}

		currentBest, exists := pr.rec[w.s]
		if !exists || w.acc <= currentBest {
			pr.rec[w.s] = w.acc
		} else {
			pr.mu.Unlock()
			break
		}

		if x := w.s[0]; x == end {
			if pr.best == -1 || w.acc <= pr.best {
				if pr.best == -1 || w.acc < pr.best {
					pr.best = w.acc
					pr.t = make(Traversed)
				}
				pr.t.merge(w.t)
			}

			pr.mu.Unlock()
			break
		}

		pr.mu.Unlock()

		candidateStates := pr.getCandidateStates(w.s)
		wg.Add(len(candidateStates) - 1)
		for _, s := range candidateStates[1:] {
			wNew := initWalker(
				w.acc+getTransitionCost(w.s, s),
				s,
				w.t,
			)

			go func() {
				pr.parallelBFS(wNew, end)
				wg.Done()
			}()
		}

		w.acc += getTransitionCost(w.s, candidateStates[0])
		w.s = candidateStates[0]
		w.t[w.s[0]] = true
	}

	wg.Wait()
}

func (m Maze) getEndpoints() (start, end Coordinate) {
	for i, row := range m {
		for j, val := range row {
			switch val {
			case 'S':
				start = Coordinate{i, j}
			case 'E':
				end = Coordinate{i, j}
			default:
				continue
			}
		}
	}
	return
}

func initParallelRecord() ParallelRecord {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	pr := ParallelRecord{
		best: -1,
		mu:   new(sync.Mutex),
		rec:  make(Record),
		m:    make(Maze, n),
		t:    nil,
	}

	for i := 0; scanner.Scan(); i++ {
		pr.m[i] = []byte(scanner.Text())
	}

	return pr
}

func main() {
	pr := initParallelRecord()

	start, end := pr.m.getEndpoints()
	pr.parallelBFS(initWalker(
		0,
		State{start, Coordinate{0, 1}},
		nil,
	), end)

	println(pr.best)
	println(len(pr.t))
}
