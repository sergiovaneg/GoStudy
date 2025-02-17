package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

const sampleSizeA = 40000000
const sampleSizeB = 40000000

const factorA = 16807
const factorB = 48271
const factorC = 2147483647

type Machine struct {
	factor      int
	denominator int
}

func (m Machine) genNext(x int) int {
	return (x * m.factor) % m.denominator
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	mA := Machine{factor: factorA, denominator: factorC}
	mB := Machine{factor: factorB, denominator: factorC}
	vals := [2]int{}

	for i := range 2 {
		scanner.Scan()
		num := regexp.MustCompile(`\d+`).FindString(scanner.Text())
		vals[i], _ = strconv.Atoi(num)
	}

	var resA int
	var qB [2][]int
	qB[0], qB[1] = make([]int, 0), make([]int, 0)
	for i := range sampleSizeA {
		vals[0] = mA.genNext(vals[0])
		vals[1] = mB.genNext(vals[1])

		if vals[0]&0xFFFF == vals[1]&0xFFFF {
			resA++
		}

		if i >= sampleSizeB {
			continue
		}

		if vals[0]&0b11 == 0 {
			qB[0] = append(qB[0], vals[0])
		}

		if vals[1]&0b111 == 0 {
			qB[1] = append(qB[1], vals[1])
		}
	}
	println(resA)

	var resB int
	for i := range min(len(qB[0]), len(qB[1])) {
		if qB[0][i]&0xFFFF == qB[1][i]&0xFFFF {
			resB++
		}
	}
	println(resB)
}
