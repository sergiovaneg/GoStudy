package main

import (
	"bufio"
	"os"
	"slices"

	"github.com/sergiovaneg/GoStudy/utils"
)

func processGroupAny(lines [][]rune, c chan<- int) {
	uniqueBuffer := make([]rune, 0)
	for _, line := range lines {
		for _, r := range line {
			uniqueBuffer, _ = utils.SortedUniqueInsert(uniqueBuffer, r)
		}
	}

	c <- len(uniqueBuffer)
}

func processGroupAll(lines [][]rune, c chan<- int) {
	uniqueBuffer := make([]rune, 0)

	for _, question := range lines[0] {
		var invalid bool
		for _, line := range lines[1:] {
			_, found := slices.BinarySearch(line, question)
			if !found {
				invalid = true
				break
			}
		}

		if !invalid {
			uniqueBuffer, _ = utils.SortedUniqueInsert(uniqueBuffer, question)
		}
	}

	c <- len(uniqueBuffer)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	buffer := make([][]rune, 0)
	var gCount int
	cA, cB := make(chan int), make(chan int)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			gCount++
			go processGroupAny(buffer, cA)
			go processGroupAll(buffer, cB)
			buffer = make([][]rune, 0)
		} else {
			runeBuffer := []rune(line)
			slices.Sort(runeBuffer)
			buffer = append(buffer, runeBuffer)
		}
	}

	var resA, resB int
	for range gCount {
		resA += <-cA
		resB += <-cB
	}
	close(cA)
	close(cB)

	println(resA)
	println(resB)
}
