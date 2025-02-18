package main

import (
	"bufio"
	"os"
	"strconv"
)

const finalSizeA = 2017 + 1
const finalSizeB = 50000000 + 1

type State struct {
	buffer      []int
	currentSize int
	idx         int
}

func getFinalState(stepSize, finalSize int) State {
	s := State{buffer: make([]int, finalSize), currentSize: 1}

	for step := range finalSize - 1 {
		s.idx = (s.idx + stepSize) % s.currentSize

		copy(s.buffer[s.idx+2:], s.buffer[s.idx+1:s.currentSize])
		s.buffer[s.idx+1] = step + 1
		s.idx++

		s.currentSize++
	}

	return s
}

func mockFinalState(stepSize, finalSize int) int {
	currentNext := -1

	for i, idx := 0, 0; i < finalSize-1; i++ {
		idx = (idx + stepSize) % (i + 1)

		if idx == 0 {
			currentNext = i + 1
		}

		idx++
	}

	return currentNext
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	stepSize, _ := strconv.Atoi(scanner.Text())

	finalState := getFinalState(stepSize, finalSizeA)
	println(finalState.buffer[(finalState.idx+1)%finalSizeA])

	println(mockFinalState(stepSize, finalSizeB))
}
