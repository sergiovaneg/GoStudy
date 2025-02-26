package main

import (
	"bufio"
	"os"

	"github.com/sergiovaneg/GoStudy/AoC/2017/fractalArt/fractals"
	"github.com/sergiovaneg/GoStudy/utils"
)

const SEED = ".#./..#/###"
const targetA = 5
const targetB = 18

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	lines := make([]string, 0, n)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	println(fractals.NaiveSolver{}.Solve(SEED, targetA, lines))
	println(fractals.NaiveSolver{}.Solve(SEED, targetB, lines))
}
