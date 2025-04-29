package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Mem map[int]int

func (m *Mem) iterate(last, turn int) int {
	var next int

	if idx, found := (*m)[last]; found {
		next = turn - idx
	} else {
		next = 0
	}
	(*m)[last] = turn

	return next
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	mem := make(Mem)

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	for idx, num := range strings.Split(scanner.Text(), ",") {
		val, _ := strconv.Atoi(num)
		mem[val] = idx
	}

	last := 0
	for turn := len(mem); turn < 2019; turn++ {
		last = mem.iterate(last, turn)
	}
	println(last)
	for turn := 2019; turn < 30000000-1; turn++ {
		last = mem.iterate(last, turn)
	}
	println(last)
}
