package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

const numPrograms = 16
const numDances = 1000000000

type State [numPrograms]byte
type Record map[State]State

func initState() State {
	var s State
	for i := range s {
		s[i] = 'a' + byte(i)
	}
	return s
}

func (s *State) updateState(inst string) {
	if inst[0] == 's' {
		idx, _ := strconv.Atoi(inst[1:])

		tmp := make([]byte, idx)
		copy(tmp, s[numPrograms-idx:])
		copy(s[idx:], s[:numPrograms-idx])
		copy(s[:idx], tmp)
	} else {
		var idxs [2]int

		if inst[0] == 'x' {
			for i, num := range strings.SplitN(inst[1:], "/", 2) {
				idxs[i], _ = strconv.Atoi(num)
			}
		} else {
			for i, val := range strings.SplitN(inst[1:], "/", 2) {
				idxs[i] = slices.Index(s[:], val[0])
			}
		}

		tmp := s[idxs[0]]
		s[idxs[0]] = s[idxs[1]]
		s[idxs[1]] = tmp
	}
}

func (r *Record) executeDance(s0 State, instList []string) State {
	if s, ok := (*r)[s0]; ok {
		return s
	}

	s := s0
	for _, inst := range instList {
		s.updateState(inst)
	}

	(*r)[s0] = s

	return s
}

func (r Record) getLoopLength(s0 State) int {
	res := 1

	for s := r[s0]; s != s0; res++ {
		s = r[s]
	}

	return res
}

func (s State) toString() string {
	str := ""
	for _, c := range s {
		str += string(c)
	}
	return str
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	s := initState()
	instList := strings.Split(scanner.Text(), ",")

	r := make(Record)
	s = r.executeDance(s, instList)

	println(s.toString())

	for i := 0; i < numDances-1; i++ {
		s = r.executeDance(s, instList)

		if _, ok := r[s]; ok {
			ll := r.getLoopLength(s)

			for range (numDances - 2 - i) % ll {
				s = r.executeDance(s, instList)
			}
			break
		}
	}

	println(s.toString())
}
