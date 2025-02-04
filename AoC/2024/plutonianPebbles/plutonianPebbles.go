package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

const nA, nB = 25, 75

type State map[int]int

func parseInput(line string) State {
	nums := strings.Split(line, " ")
	arr := make(State, len(nums))
	for _, num := range nums {
		if n, err := strconv.Atoi(num); err == nil {
			arr[n]++
		}
	}
	return arr
}

func applyRules(x int) []int {
	if x == 0 {
		return []int{1}
	}

	if l := len(strconv.Itoa(x)); l&0x01 == 0 {
		aux := utils.IPow(10, l>>1)
		return []int{x / aux, x % aux}
	}

	return []int{2024 * x}
}

func blinkN(x0 State, n int) State {
	for range n {
		x1 := make(State, len(x0))
		for pebble0 := range x0 {
			for _, pebble1 := range applyRules(pebble0) {
				x1[pebble1] += x0[pebble0]
			}
		}
		x0 = x1
	}

	return x0
}

func (x State) countPebbles() int {
	acc := 0

	for _, nPebbles := range x {
		acc += nPebbles
	}

	return acc
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	x := parseInput(scanner.Text())

	x = blinkN(x, nA)
	println(x.countPebbles())
	x = blinkN(x, nB-nA)
	println(x.countPebbles())
}
