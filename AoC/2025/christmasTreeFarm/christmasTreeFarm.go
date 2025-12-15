package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type present [3][3]bool

func parseRow(line string) [3]bool {
	var ret [3]bool

	for idx := range 3 {
		if line[idx] == '#' {
			ret[idx] = true
		}
	}

	return ret
}

func (p present) getInnerArea() int {
	acc := 0

	for _, row := range p {
		for _, filled := range row {
			if filled {
				acc++
			}
		}
	}

	return acc
}

func isValidRegion(w, h int, quantities []int, presents []present) bool {
	outerArea := 0
	for _, qty := range quantities {
		outerArea += 9 * qty
	}
	if outerArea <= w*h {
		return true
	}

	innerArea := 0
	for idx, qty := range quantities {
		innerArea += presents[idx].getInnerArea() * qty
	}
	if innerArea > w*h {
		return false
	}

	panic("Missing non-trivial implementation.")
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	presents := make([]present, 0)
	resA := 0
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ":")
		if splitLine[1] == "" {
			var p present
			for idx := range 3 {
				scanner.Scan()
				p[idx] = parseRow(scanner.Text())
			}

			presents = append(presents, p)
			scanner.Scan()
		} else {
			numbers := regexp.MustCompile(`\d+`).FindAllString(
				line, len(presents)+2,
			)
			w, _ := strconv.Atoi(numbers[0])
			h, _ := strconv.Atoi(numbers[1])
			quantities := make([]int, len(presents))
			for idx, num := range numbers[2:] {
				quantities[idx], _ = strconv.Atoi(num)
			}
			if isValidRegion(w, h, quantities, presents) {
				resA++
			}
		}
	}

	println(resA)
}
