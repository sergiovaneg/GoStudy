package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	knotHashUtils "github.com/sergiovaneg/GoStudy/AoC/2017/knotHash/knotHashUtils"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	s := knotHashUtils.InitState()

	iLengths := make([]int, 0)
	for _, num := range strings.Split(scanner.Text(), ",") {
		aux, _ := strconv.Atoi(num)
		iLengths = append(iLengths, aux)
	}
	s.SparseHash(iLengths, 1)

	println(s.Lst[0] * s.Lst[1])

	println(knotHashUtils.KnotHash(scanner.Text()))
}
