package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/sergiovaneg/GoStudy/utils"
)

func parseOperands(line string) []int {
	matches := regexp.MustCompile(`(\d+)`).FindAllString(line, -1)
	operands := make([]int, 0, len(matches))
	for _, match := range matches {
		if num, err := strconv.Atoi(match); err == nil {
			operands = append(operands, num)
		}
	}
	return operands
}

type Operator func(int, int) int

func Sum(acc, operand int) int {
	return acc + operand
}
func Prod(acc, operand int) int {
	return acc * operand
}
func Concat(acc, operand int) int {
	num, _ := strconv.Atoi(strconv.Itoa(acc) + strconv.Itoa(operand))
	return num
}

func testOperators(target, acc int, operands []int, operators []Operator) bool {
	if acc > target {
		return false
	}

	if len(operands) == 0 {
		return acc == target
	}

	for _, f := range operators {
		if testOperators(target, f(acc, operands[0]), operands[1:], operators) {
			return true
		}
	}

	return false
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)
	operandsList := make([][]int, 0, n)

	for scanner.Scan() {
		operandsList = append(operandsList, parseOperands(scanner.Text()))
	}

	res := 0
	failedIdxList := make([]int, 0, n)
	for idx, operands := range operandsList {
		if testOperators(
			operands[0], operands[1], operands[2:],
			[]Operator{Prod, Sum}) {
			res += operands[0]
		} else {
			failedIdxList = append(failedIdxList, idx)
		}
	}
	println(res)

	for _, idx := range failedIdxList {
		operands := operandsList[idx]
		if testOperators(
			operands[0], operands[1], operands[2:],
			[]Operator{Concat, Prod, Sum}) {
			res += operands[0]
		}
	}
	println(res)
}
