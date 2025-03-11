package main

import (
	"bufio"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

type Registry map[string]int

func (r Registry) get(key string) int {
	if num, err := strconv.Atoi(key); err == nil {
		return num
	} else {
		return r[key]
	}
}

func (r *Registry) execute(inst string) int {
	op := strings.SplitN(inst, " ", 3)
	switch op[0] {
	case "set":
		(*r)[op[1]] = r.get(op[2])
		return 1
	case "sub":
		(*r)[op[1]] -= r.get(op[2])
		return 1
	case "mul":
		(*r)["mulCnt"]++
		(*r)[op[1]] *= r.get(op[2])
		return 1
	case "jnz":
		if r.get(op[1]) == 0 {
			return 1
		}
		return r.get(op[2])
	}
	return 0
}

func run(instSet []string) int {
	r, idx := make(Registry), 0
	r["a"] = 0
	for idx < len(instSet) {
		idx += r.execute(instSet[idx])
	}

	return r["mulCnt"]
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	instructions := make([]string, 0, n)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	println(run(instructions))

	// Part 2, using heuristics
	lb, _ := strconv.Atoi(
		strings.SplitN(instructions[0], " ", 3)[2])
	step, _ := strconv.Atoi(
		strings.SplitN(instructions[4], " ", 3)[2])
	lb *= step
	step, _ = strconv.Atoi(
		strings.SplitN(instructions[5], " ", 3)[2])
	lb -= step

	n_iters, _ := strconv.Atoi(
		strings.SplitN(instructions[7], " ", 3)[2])
	step, _ = strconv.Atoi(
		strings.SplitN(instructions[len(instructions)-2], " ", 3)[2])
	n_iters /= step

	res := 0
	for range n_iters + 1 {
		if !big.NewInt(int64(lb)).ProbablyPrime(0) {
			res++
		}
		lb -= step
	}
	println(res)
}
