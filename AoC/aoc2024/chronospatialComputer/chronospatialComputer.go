package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Registry [3]int
type Program []int

type State struct {
	registry Registry
	program  Program
	ptr      int
	out      string
}

type InstructionSet [0o10]func(*State)

type StopCondition func(State) bool

func initState() (s State) {
	return State{
		program: make(Program, 0),
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

func (program Program) toDotfile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(f)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	getOperandRune := func(x int) string {
		if x <= 3 {
			return string('0' + rune(x))
		} else if x <= 6 {
			return string('A' + rune(x-4))
		} else {
			return "🦊"
		}
	}

	limitEndpoint := func(x int) any {
		if x >= len(program) {
			return "end"
		}
		return x
	}

	fmt.Fprintln(w, "strict digraph{")
	fmt.Fprintln(w, "\tlayout=\"sfdp\";")
	fmt.Fprintln(w, "\tStart -> 0;")

	for i := 0; i < len(program); i += 2 {
		switch program[i] {
		case 0:
			fmt.Fprintf(w, "\t%v -> %v [label=\"A >>= %v\"];\n",
				i, limitEndpoint(i+2), getOperandRune(program[i+1]))
		case 1:
			fmt.Fprintf(w, "\t%v -> %v [label=\"B ^= %v\"];\n",
				i, limitEndpoint(i+2), program[i+1])
		case 2:
			fmt.Fprintf(w, "\t%v -> %v [label=\"B = %v & 0o7\"];\n",
				i, limitEndpoint(i+2), getOperandRune(program[i+1]))
		case 3:
			fmt.Fprintf(w, "\t%v -> %v [label=\"if A == 0\"];\n",
				i, limitEndpoint(i+2))
			fmt.Fprintf(w, "\t%v -> %v [label=\"if A != 0\"];\n",
				i, limitEndpoint(program[i+1]))
		case 4:
			fmt.Fprintf(w, "\t%v -> %v [label=\"B ^= C\"];\n", i, limitEndpoint(i+2))
		case 5:
			fmt.Fprintf(w, "\t%v -> %v [label=\"print(%v & 0o7)\"];\n",
				i, limitEndpoint(i+2), getOperandRune(program[i+1]))
		case 6:
			fmt.Fprintf(w, "\t%v -> %v [label=\"B = A >> %v\"];\n",
				i, limitEndpoint(i+2), getOperandRune(program[i+1]))
		case 7:
			fmt.Fprintf(w, "\t%v -> %v [label=\"C = A >> %v\"];\n",
				i, limitEndpoint(i+2), getOperandRune(program[i+1]))
		}
	}

	fmt.Fprintln(w, "}")
}

func (p Program) recursiveSolve(level int, solution *int) bool {
	bTarget, shift := p[level], level<<1+level

	aRelevant := ((*solution) >> shift) & ^(0b111)

	for aLevel := 0; aLevel < 8; aLevel++ {
		aTest := aRelevant ^ aLevel
		aTest = aLevel ^ 0b001 ^ (aTest >> (aLevel ^ 0b010))
		if aTest&0b111 == bTarget {
			*solution &= ^((0b111 ^ aLevel) << shift)
			*solution |= aLevel << shift
			if level == 0 || p.recursiveSolve(level-1, solution) {
				return true
			}
		}
	}

	return false
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

	s.program.toDotfile("./program.dot")

	var aOpt int
	if f := s.program.recursiveSolve(len(s.program)-1, &aOpt); f {
		println(aOpt)
	}
}
