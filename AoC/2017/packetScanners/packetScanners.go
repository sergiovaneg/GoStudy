package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

type Layer [2]int

func parseLayer(line string) Layer {
	var l Layer
	for i, num := range strings.SplitN(line, ": ", 2) {
		l[i], _ = strconv.Atoi(num)
	}
	return l
}

func (l Layer) period() int {
	return (l[1] - 1) << 1
}

func getSeverityScore(layers []Layer, delay int) int {
	var res int

	for _, layer := range layers {
		if (layer[0]+delay)%layer.period() == 0 {
			res += layer[0] * layer[1]
		}
	}

	return res
}

func isCaught(layers []Layer, delay int) bool {
	for _, layer := range layers {
		if (layer[0]+delay)%layer.period() == 0 {
			return true
		}
	}

	return false
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	layers := make([]Layer, 0, n)

	for scanner.Scan() {
		layers = append(layers, parseLayer(scanner.Text()))
	}

	resA := getSeverityScore(layers, 0)
	println(resA)

	var resB int
	for isCaught(layers, resB) {
		resB++
	}
	println(resB)
}
