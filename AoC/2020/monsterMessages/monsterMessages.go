package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	isEndpoint  bool
	endpointStr string
	nextRules   [][]int
}

type CallArgs struct {
	msg string
	id  int
}

type Ruleset struct {
	rules  map[int]Rule
	memory map[CallArgs]bool
}

func parseRule(line string) (int, Rule) {
	subStrs := strings.Split(line, ": ")
	id, _ := strconv.Atoi(subStrs[0])

	if strings.Contains(subStrs[1], "\"") {
		return id, Rule{isEndpoint: true, endpointStr: subStrs[1][1:2]}
	}

	paths := make([][]int, 0)
	for pathStr := range strings.SplitSeq(subStrs[1], " | ") {
		path := make([]int, 0)
		for nextStr := range strings.SplitSeq(pathStr, " ") {
			next, _ := strconv.Atoi(nextStr)
			path = append(path, next)
		}
		paths = append(paths, path)
	}

	return id, Rule{isEndpoint: false, nextRules: paths}
}

func partitionString(s string, k int) [][]string {
	if k == 1 {
		return [][]string{{s}}
	}

	partitions := make([][]string, 0)
	for idx := range s {
		lhs := s[:idx]
		for _, rhs := range partitionString(s[idx:], k-1) {
			partitions = append(partitions, append([]string{lhs}, rhs...))
		}
	}

	return partitions
}

func (rs *Ruleset) query(msg string, id int) bool {
	callArgs := CallArgs{msg, id}
	if res, ok := rs.memory[callArgs]; ok {
		return res
	}

	rs.memory[callArgs] = rs.validate(msg, id)
	return rs.memory[callArgs]
}

func (rs *Ruleset) validate(msg string, id int) bool {
	if msg == "" {
		return false
	}

	if rs.rules[id].isEndpoint {
		return msg == rs.rules[id].endpointStr
	}

	for _, path := range rs.rules[id].nextRules {
		for _, part := range partitionString(msg, len(path)) {
			valid := true
			for idx := range part {
				valid = valid && rs.query(part[idx], path[idx])
			}
			if valid {
				return true
			}
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
	ruleset := Ruleset{
		rules:  make(map[int]Rule),
		memory: make(map[CallArgs]bool),
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		id, rule := parseRule(line)
		ruleset.rules[id] = rule
	}

	messages := make([]string, 0)
	for scanner.Scan() {
		messages = append(messages, scanner.Text())
	}

	resA := 0
	for _, msg := range messages {
		if ruleset.validate(msg, 0) {
			resA++
		}
	}
	println(resA)

	ruleset.memory = make(map[CallArgs]bool)
	ruleset.rules[8] = Rule{
		isEndpoint: false,
		nextRules:  [][]int{{42}, {42, 8}},
	}
	ruleset.rules[11] = Rule{
		isEndpoint: false,
		nextRules:  [][]int{{42, 31}, {42, 11, 31}},
	}

	resB := 0
	for _, msg := range messages {
		if ruleset.validate(msg, 0) {
			resB++
		}
	}
	println(resB)
}
