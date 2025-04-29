package main

import (
	"bufio"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type rangePair [2][2]int
type ruleset map[string]rangePair
type ticket []int

func parseRule(line string) (string, rangePair) {
	split := strings.SplitN(line, ": ", 2)
	var ranges rangePair
	for idx, num := range regexp.MustCompile(`\d+`).FindAllString(split[1], -1) {
		ranges[idx>>1][idx&0x01], _ = strconv.Atoi(num)
	}

	return split[0], ranges
}

func parseTicket(line string) ticket {
	nums := strings.Split(line, ",")
	t := make(ticket, len(nums))

	for idx, num := range nums {
		t[idx], _ = strconv.Atoi(num)
	}

	return t
}

func (rp rangePair) isWithin(x int) bool {
	return (x >= rp[0][0] && x <= rp[0][1]) || (x >= rp[1][0] && x <= rp[1][1])
}

func (r ruleset) getErrorRate(t ticket) int {
	score := 0

	for _, val := range t {
		var flag bool
		for _, rPair := range r {
			if rPair.isWithin(val) {
				flag = true
				break
			}
		}
		if !flag {
			score += val
		}
	}

	return score
}

func (rp rangePair) isValidAssignment(tickets []ticket, colIdx int) bool {
	for _, t := range tickets {
		if !rp.isWithin(t[colIdx]) {
			return false
		}
	}
	return true
}

func (r ruleset) getSortedFields(tickets []ticket) []string {
	valid, auxRange := make(map[int][]string), make([]int, 0, len(r))

	for colIdx := range len(r) {
		auxRange = append(auxRange, colIdx)
		fields := make([]string, 0)
		for field, rp := range r {
			if rp.isValidAssignment(tickets, colIdx) {
				fields = append(fields, field)
			}
		}
		valid[colIdx] = fields
	}
	slices.SortFunc(auxRange, func(a, b int) int {
		return len(valid[a]) - len(valid[b])
	})

	consumed := make(map[string]bool, len(r))
	sortedKeys := make([]string, len(r))
	for _, colIdx := range auxRange {
		for _, field := range valid[colIdx] {
			if consumed[field] {
				continue
			}
			sortedKeys[colIdx] = field
			consumed[field] = true
			break
		}
	}

	return sortedKeys
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	r := make(ruleset)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		field, ranges := parseRule(line)
		r[field] = ranges
	}

	tickets := make([]ticket, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if !regexp.MustCompile(`^\d`).MatchString(line) {
			continue
		}
		tickets = append(tickets, parseTicket(line))
	}

	var resA int
	for _, t := range tickets[1:] {
		resA += r.getErrorRate(t)
	}
	println(resA)

	tickets = slices.DeleteFunc(tickets, func(t ticket) bool {
		return r.getErrorRate(t) > 0
	})

	sortedFields := r.getSortedFields(tickets[1:])
	resB := 1
	for colIdx, field := range sortedFields {
		if !strings.HasPrefix(field, "departure") {
			continue
		}
		resB *= tickets[0][colIdx]
	}
	println(resB)
}
