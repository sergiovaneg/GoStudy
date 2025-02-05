package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const blockCount = 16

type State [blockCount]int
type Record map[State]bool

func (r *Record) conditionalUpdate(s State) bool {
	if _, ok := (*r)[s]; ok {
		return false
	}

	(*r)[s] = true
	return true
}

func (s *State) mutate() {
	i, ref := 0, s[0]

	for j, val := range s {
		if val > ref {
			i = j
			ref = val
		}
	}

	s[i] = 0

	i++
	q, r := ref/blockCount, ref%blockCount
	for j := range r {
		s[(i+j)%blockCount] += q + 1
	}

	i += r
	for j := range blockCount - r {
		s[(i+j)%blockCount] += q
	}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var s State

	for i, num := range strings.Split(scanner.Text(), "\t") {
		s[i], _ = strconv.Atoi(num)
	}

	var resA int
	for r := make(Record); r.conditionalUpdate(s); s.mutate() {
		resA++
	}
	println(resA)

	var resB int
	for r := make(Record); r.conditionalUpdate(s); s.mutate() {
		resB++
	}
	println(resB)
}
