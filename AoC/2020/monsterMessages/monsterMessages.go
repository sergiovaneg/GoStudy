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

type MemKey struct {
	id  int
	msg string
}
type Memory map[MemKey][][2]int
type Ruleset struct {
	rules map[int]func(string) [][2]int
	dp    Memory
}

func isContiguous(rangeArr [][2]int) bool {
	for i, r1 := range rangeArr[1:] {
		if rangeArr[i][1] != r1[0] {
			return false
		}
	}
	return true
}

func initRuleset() Ruleset {
	return Ruleset{
		rules: make(map[int]func(string) [][2]int),
		dp:    make(Memory),
	}
}

func (rs *Ruleset) call(id int, msg string) [][2]int {
	key := MemKey{id, msg}
	if res, ok := rs.dp[key]; ok {
		return res
	} else {
		res = rs.rules[id](msg)
		rs.dp[key] = slices.Clone(res)
		return res
	}
}

func (rs *Ruleset) addRule(line string) {
	substr := strings.Split(line, ": ")
	id, _ := strconv.Atoi(substr[0])

	if j := strings.Index(substr[1], `"`); j != -1 {
		re := regexp.MustCompile(substr[1][j+1 : j+2])
		rs.rules[id] = func(msg string) [][2]int {
			candidates := make([][2]int, 0)
			for _, idxs := range re.FindAllStringIndex(msg, -1) {
				candidates = append(candidates, [2]int{idxs[0], idxs[1]})
			}
			return candidates
		}
	} else {
		options := make([][]int, 0)
		for i, opt := range strings.Split(substr[1], " | ") {
			options = append(options, make([]int, 0))
			for _, num := range regexp.MustCompile(
				`\d+`).FindAllString(opt, -1) {
				val, _ := strconv.Atoi(num)
				options[i] = append(options[i], val)
			}
		}

		rs.rules[id] = func(msg string) [][2]int {
			candidates := make([][2]int, 0)

			for _, opt := range options {
				cardinality := len(opt)
				pools := make([][][2]int, cardinality)
				for poolIdx, ruleIdx := range opt {
					pools[poolIdx] = rs.call(ruleIdx, msg)
				}

				for _, seq := range utils.Choices(pools) {
					if !isContiguous(seq) {
						continue
					}
					candidates = append(candidates, [2]int{
						seq[0][0], seq[cardinality-1][1]})
				}
			}

			return candidates
		}
	}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rs := initRuleset()

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		rs.addRule(line)
	}

	resA := 0
	for scanner.Scan() {
		msg, flag := scanner.Text(), false
		candidates := rs.call(0, msg)
		for _, cand := range candidates {
			match := msg[cand[0]:cand[1]]
			if match == msg {
				flag = true
				break
			}
		}
		if flag {
			resA++
		}
	}
	println(resA)
}
