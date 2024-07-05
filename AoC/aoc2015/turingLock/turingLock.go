package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

type Memory map[byte]uint

func (m Memory) run(program []string) {
	n := len(program)
	for idx := 0; idx < n && idx >= 0; {
		line := program[idx]
		split_idx := strings.Index(line, " ")
		inst := line[:split_idx]
		switch inst {
		case "hlf":
			m[line[split_idx+1]] >>= 1
			idx++
		case "tpl":
			m[line[split_idx+1]] += m[line[split_idx+1]] << 1
			idx++
		case "inc":
			m[line[split_idx+1]]++
			idx++
		case "jmp":
			offset, _ := strconv.Atoi(line[split_idx+1:])
			idx += offset
		case "jie":
			if m[line[split_idx+1]]&0x01 == 0x00 {
				offset, _ := strconv.Atoi(line[split_idx+4:])
				idx += offset
			} else {
				idx++
			}
		case "jio":
			if m[line[split_idx+1]] == 1 {
				offset, _ := strconv.Atoi(line[split_idx+4:])
				idx += offset
			} else {
				idx++
			}
		}
	}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	n, _ := utils.LineCounter(file)
	program := make([]string, 0, n)
	memory := Memory{
		'a': 0,
		'b': 0,
	}

	for scanner.Scan() {
		program = append(program, scanner.Text())
	}

	memory.run(program)
	println(memory['b'])

	memory['a'], memory['b'] = 1, 0
	memory.run(program)
	println(memory['b'])
}
