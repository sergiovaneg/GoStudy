package main

import (
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Turing struct {
	tape    map[int]int
	cursor  int
	state   string
	instSet map[string]func()
}

const stateRe = `(?:state) ([A-Z]{1})`

func getParams(inst string) (string, int) {
	initial := regexp.MustCompile(stateRe).FindStringSubmatch(inst)[1]
	numIters, _ := strconv.Atoi(
		regexp.MustCompile(`\d+`).FindString(inst))
	return initial, numIters
}

func (t *Turing) addInstruction(inst string) {
	lines := strings.Split(inst, "\n")
	src := regexp.MustCompile(stateRe).FindStringSubmatch(lines[0])[1]

	var w [2]int
	var delta [2]int
	var next [2]string

	for i := range 2 {
		if strings.Contains(lines[2+i<<2], "1") {
			w[i] = 1
		} else {
			w[i] = 0
		}

		if strings.Contains(lines[3+i<<2], "right") {
			delta[i] = 1
		} else {
			delta[i] = -1
		}

		next[i] = regexp.MustCompile(stateRe).FindStringSubmatch(
			lines[4+i<<2])[1]
	}

	t.instSet[src] = func() {
		r := t.tape[t.cursor]
		t.tape[t.cursor] = w[r]
		t.cursor += delta[r]
		t.state = next[r]
	}
}

func (t Turing) getChecksum() int {
	var res int

	for _, v := range t.tape {
		res += v
	}

	return res
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buffer, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	program := strings.Split(string(buffer), "\n\n")
	initialState, numIters := getParams(program[0])
	turing := Turing{
		state:   initialState,
		tape:    make(map[int]int),
		instSet: make(map[string]func()),
	}

	for _, inst := range program[1:] {
		turing.addInstruction(inst)
	}

	for range numIters {
		turing.instSet[turing.state]()
	}

	println(turing.getChecksum())
}
