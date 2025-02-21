package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

func recoverFrequency(instructions []string) int {
	reg := make(map[string]int)
	ptr, n := 0, len(instructions)

	for ptr >= 0 && ptr < n {
		inst := strings.Split(instructions[ptr], " ")
		aux, err := strconv.Atoi(inst[len(inst)-1])
		if err != nil {
			aux = reg[inst[len(inst)-1]]
		}

		switch strings.ToLower(inst[0]) {
		case "snd":
			reg[""] = aux
		case "set":
			reg[inst[1]] = aux
		case "add":
			reg[inst[1]] += aux
		case "mul":
			reg[inst[1]] *= aux
		case "mod":
			reg[inst[1]] %= aux
		case "rcv":
			if aux > 0 {
				ptr = n //break
			}
		case "jgz":
			crit, err := strconv.Atoi(inst[1])
			if err != nil {
				crit = reg[inst[1]]
			}

			if crit > 0 {
				ptr += aux - 1
			}
		}

		ptr++
	}

	return reg[""]
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	instructions := make([]string, n)

	for idx := 0; scanner.Scan(); idx++ {
		instructions[idx] = scanner.Text()
	}

	println(recoverFrequency(instructions))
}
