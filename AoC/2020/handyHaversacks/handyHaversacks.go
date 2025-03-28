package main

import (
	"bufio"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

type Ruleset map[string]map[string]int

func (r *Ruleset) update(line string) {
	matches := regexp.MustCompile(
		`(\d+ )*(\w+\s\w+)(?:\sbag)`).FindAllString(line, -1)
	for idx, match := range matches {
		matches[idx], _ = strings.CutSuffix(match, " bag")
	}

	(*r)[matches[0]] = make(map[string]int)
	if matches[1] == "no other" {
		return
	}

	for _, match := range matches[1:] {
		idxs := regexp.MustCompile(`\d+`).FindStringIndex(match)
		cap, _ := strconv.Atoi(match[:idxs[1]])
		(*r)[matches[0]][match[idxs[1]+1:]] = cap
	}
}

func (r Ruleset) outerPopulate(current string, explored *[]string) {
	for outer, innerSet := range r {
		if slices.Contains(*explored, outer) {
			continue
		}

		var contains bool
		for inner := range innerSet {
			if inner == current {
				contains = true
				break
			}
		}

		if contains {
			*explored = append(*explored, outer)
			r.outerPopulate(outer, explored)
		}
	}
}

func (r Ruleset) innerCount(current string, dp *map[string]int) int {
	if dp == nil {
		dp = new(map[string]int)
		*dp = make(map[string]int)
	} else {
		if val, ok := (*dp)[current]; ok {
			return val
		}
	}

	res := 1
	for key, value := range r[current] {
		res += value * r.innerCount(key, dp)
	}

	(*dp)[current] = res
	return res
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	r := make(Ruleset, n)

	for scanner.Scan() {
		r.update(scanner.Text())
	}

	outer := make([]string, 0, n)
	r.outerPopulate("shiny gold", &outer)
	println(len(outer))

	println(r.innerCount("shiny gold", nil) - 1)
}
