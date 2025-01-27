package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Registry [3]int

type State struct {
	registry Registry
	program  []int
	ptr      int
	out      string
}

type InstructionSet [0o10]func(*State)

type StopCondition func(State) bool

func initState() (s State) {
	return State{
		program: make([]int, 0),
	}
}

func (s State) getComboOperand() int {
	x := s.program[s.ptr+1]
	if x <= 3 {
		return x
	} else if x <= 6 {
		return s.registry[x-4]
	} else {
		log.Fatal("Invalid operand.")
		return -1
	}
}

func initInstructionSet() InstructionSet {
	var is InstructionSet

	is[0] = func(s *State) {
		s.registry[0] >>= s.getComboOperand()
		s.ptr += 2
	}

	is[1] = func(s *State) {
		s.registry[1] ^= s.program[s.ptr+1]
		s.ptr += 2
	}

	is[2] = func(s *State) {
		s.registry[1] = s.getComboOperand() & 0o7
		s.ptr += 2
	}

	is[3] = func(s *State) {
		if s.registry[0] == 0 {
			s.ptr += 2
		} else {
			s.ptr = s.program[s.ptr+1]
		}
	}

	is[4] = func(s *State) {
		s.registry[1] ^= s.registry[2]
		s.ptr += 2
	}

	is[5] = func(s *State) {
		s.out += strconv.Itoa(s.getComboOperand()&0o7) + ","
		s.ptr += 2
	}

	is[6] = func(s *State) {
		s.registry[1] = s.registry[0] >> s.getComboOperand()
		s.ptr += 2
	}

	is[7] = func(s *State) {
		s.registry[2] = s.registry[0] >> s.getComboOperand()
		s.ptr += 2
	}

	return is
}

func (is InstructionSet) execute(s State, sc StopCondition) string {
	if sc == nil {
		sc = func(_ State) bool { return false }
	}
	s.ptr = 0
	s.out = ""
	for s.ptr < len(s.program) {
		if sc(s) {
			break
		}
		is[s.program[s.ptr]](&s)
	}

	return s.out[:len(s.out)-1]
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var r Registry

	for i := 0; scanner.Scan(); i++ {
		if scanner.Text() == "" {
			break
		}

		num, _ := strconv.Atoi(strings.Split(scanner.Text(), ": ")[1])
		r[i] = num
	}

	scanner.Scan()
	s := initState()

	programStr := strings.Split(scanner.Text(), ": ")[1]
	s.registry = r
	for i := 0; i < len(programStr); i += 2 {
		s.program = append(s.program, int(programStr[i]-'0'))
	}

	is := initInstructionSet()
	println(is.execute(s, nil))

	lb := 0
	for {
		lb++
		s.registry[0] = lb
		out := is.execute(s,
			func(s State) bool {
				l1, l2 := len(programStr), len(s.out)
				if l1 > l2 {
					return !strings.HasPrefix(programStr, s.out)
				} else {
					return l2 > l1+1
				}
			})

		if out == programStr {
			println(lb)
			break
		}
	}
}
