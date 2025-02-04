package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

func processSegment(line string) int {
	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

	result := 0
	for _, match := range re.FindAllStringSubmatch(line, -1) {
		num1, err1 := strconv.Atoi(match[1])
		num2, err2 := strconv.Atoi(match[2])
		if err1 == nil && err2 == nil {
			result += num1 * num2
		}
	}

	return result
}

func conditionalFind(line string, state bool) int {
	if state {
		return strings.Index(line, "don't()")
	} else {
		return strings.Index(line, "do()")
	}
}

func processSegmentStateful(line string, state bool) (int, bool) {
	lb, res := 0, 0
	for {
		ub := conditionalFind(line[lb:], state)
		if ub == -1 {
			if state {
				res += processSegment(line[lb:])
			}
			break
		} else {
			if state {
				res += processSegment(line[lb : lb+ub])
			}
			state = !state
			lb += ub
		}
	}

	return res, state
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	lines := make([]string, 0, n)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	res_0 := 0
	for _, line := range lines {
		res_0 += processSegment(line)
	}
	println(res_0)

	res_1, aux, state := 0, 0, true
	for _, line := range lines {
		aux, state = processSegmentStateful(line, state)
		res_1 += aux
	}
	println(res_1)
}
