package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type PathRecord map[[2]int]int

func (record *PathRecord) updatePosition(x0, v0 [2]int, inst string) (x, v [2]int) {
	n, _ := strconv.Atoi(inst[1:])
	switch inst[0] {
	case 'L':
		v[0] = -v0[1]
		v[1] = v0[0]
	case 'R':
		v[0] = v0[1]
		v[1] = -v0[0]
	}

	x = x0
	for range n {
		x[0], x[1] = x[0]+v[0], x[1]+v[1]
		(*record)[x]++
		if (*record)[x] > 1 {
			break
		}
	}

	return
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	x, v := [2]int{0, 0}, [2]int{-1, 0}
	instructions := strings.Split(scanner.Text(), ", ")

	visited := make(PathRecord, len(instructions)+1)
	visited[x] = 1

	for _, inst := range instructions {
		x, v = visited.updatePosition(x, v, inst)
		if visited[x] > 1 {
			break
		}
	}

	if x[0] < 0 {
		x[0] = -x[0]
	}
	if x[1] < 0 {
		x[1] = -x[1]
	}
	println(x[0] + x[1])
}
