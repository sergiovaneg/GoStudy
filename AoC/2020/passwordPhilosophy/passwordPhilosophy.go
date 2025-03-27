package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Policy func(string, string, int, int) bool

func policyA(str, substr string, lb, ub int) bool {
	cnt := strings.Count(str, substr)

	return cnt >= lb && cnt <= ub
}

func policyB(str, substr string, i, j int) bool {
	f1 := str[i-1:i] == substr
	f2 := str[j-1:j] == substr

	return f1 != f2
}

func isValidPassword(passwd string, policy Policy) bool {
	elems := regexp.MustCompile(`[-:\s]{1,2}`).Split(passwd, 4)
	lb, _ := strconv.Atoi(elems[0])
	ub, _ := strconv.Atoi(elems[1])

	return policy(elems[3], elems[2], lb, ub)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var resA, resB int
	for scanner.Scan() {
		passwd := scanner.Text()
		if isValidPassword(passwd, policyA) {
			resA++
		}
		if isValidPassword(passwd, policyB) {
			resB++
		}
	}

	println(resA)
	println(resB)
}
