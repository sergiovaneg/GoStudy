package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

type Inst struct {
	op  string
	arg int
}
type Program []Inst

func parseLine(line string) Inst {
	elems := strings.SplitN(line, " ", 2)
	inst := Inst{op: elems[0]}
	inst.arg, _ = strconv.Atoi(elems[1])

	return inst
}

func isolatedRun(program Program) (int, bool) {
	var acc, ptr int
	var flag bool

	n := len(program)

	prev := make([]bool, n)

	for ptr < n {
		if prev[ptr] {
			flag = true
			break
		}
		prev[ptr] = true

		inst := program[ptr]
		if inst.op == "jmp" {
			ptr += inst.arg
			continue
		}

		if inst.op == "acc" {
			acc += inst.arg
		}
		ptr++
	}

	return acc, flag
}

func mendRun(prog Program) int {
	var s0, s1 string

	for idx, inst := range prog {
		if inst.op == "acc" {
			continue
		}

		if inst.op == "jmp" {
			s0, s1 = "jmp", "nop"
		} else {
			s0, s1 = "nop", "jmp"
		}

		prog[idx].op = s1
		acc, flag := isolatedRun(prog)
		prog[idx].op = s0

		if !flag {
			return acc
		}
	}

	return -1
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	prog := make(Program, 0, n)

	for scanner.Scan() {
		prog = append(prog, parseLine(scanner.Text()))
	}

	println(isolatedRun(prog))
	println(mendRun(prog))
}
