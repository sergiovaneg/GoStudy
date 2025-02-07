package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const listSize = 256

type State struct {
	lst [listSize]int
	pos int
	skp int
}

func initState() State {
	var s State
	for i := range listSize {
		s.lst[i] = i
	}

	return s
}

func (s *State) iter(l int) {
	aux, j := make([]int, l), s.pos

	for i := range l {
		aux[i] = s.lst[j]

		if j == listSize-1 {
			j = 0
		} else {
			j++
		}
	}

	for i := range l {
		if j == 0 {
			j = listSize - 1
		} else {
			j--
		}

		s.lst[j] = aux[i]
	}

	s.pos = (s.pos + l + s.skp) % listSize
	s.skp++
}

func (s *State) sparseHash(iLengths []int, nRounds int) {
	for range nRounds {
		for _, l := range iLengths {
			s.iter(l)
		}
	}
}

func (s State) denseHash() string {
	hash := ""

	for i := 0; i < listSize; i += 16 {
		var aux int
		for j := range 16 {
			aux ^= s.lst[i+j]
		}
		hash += fmt.Sprintf("%0.2x", aux)
	}

	return hash
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

	iLengths := make([]int, 0)
	for _, num := range strings.Split(scanner.Text(), ",") {
		aux, _ := strconv.Atoi(num)
		iLengths = append(iLengths, aux)
	}
	s.sparseHash(iLengths, 1)

	println(s.lst[0] * s.lst[1])

	s = initState()

	iLengths = make([]int, 0)
	for _, c := range scanner.Text() {
		iLengths = append(iLengths, int(c))
	}
	for _, v := range [5]int{17, 31, 73, 47, 23} {
		iLengths = append(iLengths, v)
	}
	s.sparseHash(iLengths, 64)

	println(s.denseHash())
}
