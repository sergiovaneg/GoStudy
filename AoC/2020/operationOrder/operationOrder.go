package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

func operateBasic(expr string) int {
	numbers := regexp.MustCompile(`\d+`).FindAllString(expr, -1)
	ops := regexp.MustCompile(`[\+|\*]{1}`).FindAllString(expr, -1)

	res, _ := strconv.Atoi(numbers[0])
	for idx, op := range ops {
		aux, _ := strconv.Atoi(numbers[idx+1])
		if op == "*" {
			res *= aux
		} else {
			res += aux
		}
	}

	return res
}

func operateAdvanced(expr string) int {
	res := 1

	for factor := range strings.SplitSeq(expr, " * ") {
		var acc int
		for _, num := range regexp.MustCompile(`\d+`).FindAllString(factor, -1) {
			val, _ := strconv.Atoi(num)
			acc += val
		}
		res *= acc
	}

	return res
}

func eval(expr string, reducer func(string) int) int {
	groups := make([]string, 0)

	level, lb := 0, -1
	for idx, r := range expr {
		if r == '(' {
			if level == 0 {
				lb = idx
			}
			level++
		} else if r == ')' {
			level--
			if level == 0 {
				groups = append(groups, expr[lb:idx+1])
			}
		}
	}

	for _, group := range groups {
		n := len(group)
		expr = strings.ReplaceAll(
			expr, group, strconv.Itoa(eval(group[1:n-1], reducer)))
	}

	return reducer(expr)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	c := make(chan [2]int, n)

	for scanner.Scan() {
		go func(line string) {
			c <- [2]int{
				eval(line, operateBasic),
				eval(line, operateAdvanced),
			}
		}(scanner.Text())
	}

	var resA, resB int
	for range n {
		res := <-c
		resA += res[0]
		resB += res[1]
	}
	println(resA)
	println(resB)
}
