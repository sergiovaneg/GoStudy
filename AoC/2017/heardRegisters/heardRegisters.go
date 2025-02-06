package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Registry map[string]int
type State struct {
	reg Registry
	max int
}

func (s *State) run(instruction string) {
	elements := strings.SplitN(instruction, " ", 7)
	loCond := s.reg[elements[4]]
	roCond, err := strconv.Atoi(elements[6])

	if err != nil {
		panic(err)
	}

	var cond bool
	switch elements[5] {
	case ">":
		cond = loCond > roCond
	case ">=":
		cond = loCond >= roCond
	case "==":
		cond = loCond == roCond
	case "!=":
		cond = loCond != roCond
	case "<=":
		cond = loCond <= roCond
	case "<":
		cond = loCond < roCond
	}
	if !cond {
		return
	}

	delta, _ := strconv.Atoi(elements[2])
	if elements[1] == "inc" {
		s.reg[elements[0]] += delta
	} else {
		s.reg[elements[0]] -= delta
	}

	s.max = max(s.max, s.reg[elements[0]])
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var s State
	s.reg = make(Registry)

	for scanner.Scan() {
		s.run(scanner.Text())
	}

	resA := 0
	for _, v := range s.reg {
		resA = max(resA, v)
	}

	println(resA)
	println(s.max)
}
