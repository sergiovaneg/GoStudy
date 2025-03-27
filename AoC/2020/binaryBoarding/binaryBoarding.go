package main

import (
	"bufio"
	"os"
	"slices"

	"github.com/sergiovaneg/GoStudy/utils"
)

func getSeatId(serial string) int {
	var id int

	for _, c := range serial {
		id <<= 1
		if c == 'B' || c == 'R' {
			id++
		}
	}

	return id
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	ids := make([]int, 0, n)
	for scanner.Scan() {
		ids = append(ids, getSeatId(scanner.Text()))
	}
	slices.Sort(ids)

	println(ids[n-1])
	for idx, id := range ids[1:] {
		if id-ids[idx] == 2 {
			println(id - 1)
			break
		}
	}
}
