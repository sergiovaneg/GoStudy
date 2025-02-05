package main

import (
	"bufio"
	"os"
	"strconv"

	"github.com/sergiovaneg/GoStudy/utils"
)

type CPU struct {
	ptr, cnt     int
	instructions []int
}

func initCPU(inst []int) CPU {
	cpu := CPU{instructions: make([]int, len(inst))}
	copy(cpu.instructions, inst)
	return cpu
}

func (cpu *CPU) iterateA() bool {
	jmp := cpu.instructions[cpu.ptr]
	cpu.instructions[cpu.ptr]++
	cpu.ptr += jmp
	cpu.cnt++

	return cpu.ptr >= 0 && cpu.ptr < len(cpu.instructions)
}

func (cpu *CPU) iterateB() bool {
	jmp := cpu.instructions[cpu.ptr]
	if jmp >= 3 {
		cpu.instructions[cpu.ptr]--
	} else {
		cpu.instructions[cpu.ptr]++
	}
	cpu.ptr += jmp
	cpu.cnt++

	return cpu.ptr >= 0 && cpu.ptr < len(cpu.instructions)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	instructions := make([]int, n)
	for i := 0; scanner.Scan(); i++ {
		aux, _ := strconv.Atoi(scanner.Text())
		instructions[i] = aux
	}

	cpu := initCPU(instructions)
	for cpu.iterateA() {
	}
	println(cpu.cnt)

	cpu = initCPU(instructions)
	for cpu.iterateB() {
	}
	println(cpu.cnt)
}
