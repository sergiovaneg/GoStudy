package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"

	"github.com/sergiovaneg/GoStudy/utils"
)

type RuleSet map[int][]int

func (r *RuleSet) update(rule string) {
	v, vErr := strconv.Atoi(rule[:2])
	k, kErr := strconv.Atoi(rule[3:])
	if kErr == nil && vErr == nil {
		if current := (*r)[k]; current == nil {
			(*r)[k] = []int{v}
		} else {
			(*r)[k] = append(current, v)
		}

	}
}

func (r RuleSet) compare(k1, k2 int) int {
	if slices.Contains(r[k2], k1) {
		return -1
	}
	if slices.Contains(r[k1], k2) {
		return 1
	}
	return 0
}

func parseUpdate(update string) []int {
	l := len(update)
	res := make([]int, 0, (l+1)/3)
	for idx := 0; idx < l; idx += 3 {
		if v, err := strconv.Atoi(update[idx : idx+2]); err == nil {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	var i int
	r := make(RuleSet, n)
	for i = 0; scanner.Scan(); i++ {
		if scanner.Text() == "" {
			break
		}
		r.update(scanner.Text())
	}

	updates := make([][]int, 0, n-i)
	for scanner.Scan() {
		updates = append(updates, parseUpdate(scanner.Text()))
	}

	res_0, res_1 := 0, 0
	for _, update := range updates {
		if slices.IsSortedFunc(update, r.compare) {
			res_0 += update[len(update)/2]
		} else {
			slices.SortFunc(update, r.compare)
			res_1 += update[len(update)/2]
		}
	}
	println(res_0)
	println(res_1)
}
