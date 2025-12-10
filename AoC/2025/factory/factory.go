package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type machine struct {
	target      int
	buttons     []int
	numCounters int
	joltages    string
}

func parseMachine(line string) machine {
	var ret machine

	target := regexp.MustCompile(`(\.|\#)+`).FindString(line)
	for i, r := range target {
		if r == '#' {
			ret.target += 0x01 << i
		}
	}

	ret.buttons = make([]int, 0)
	buttons := regexp.MustCompile(`\([\d\,]+\)`).FindAllString(line, -1)
	for _, button := range buttons {
		mask := 0
		for num := range strings.SplitSeq(button[1:len(button)-1], ",") {
			shift, _ := strconv.Atoi(num)
			mask += 1 << shift
		}
		ret.buttons = append(ret.buttons, mask)
	}

	joltages := regexp.MustCompile(`\{[\d\,]+\}`).FindString(line)
	ret.joltages = joltages[1 : len(joltages)-1]
	ret.numCounters = len(target)

	return ret
}

func (m machine) minPushes() int {
	dp := map[int]int{0: 0}
	states := []int{0}
	cnt := 0

	for len(states) > 0 {
		newStates := make([]int, 0)
		cnt++

		for _, s0 := range states {
			for _, button := range m.buttons {
				s1 := s0 ^ button

				if _, ok := dp[s1]; !ok {
					newStates = append(newStates, s1)
					dp[s1] = cnt
				}
			}
		}

		if _, ok := dp[m.target]; ok {
			break
		}

		states = newStates
	}

	return dp[m.target]
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var resA int
	for scanner.Scan() {
		resA += parseMachine(scanner.Text()).minPushes()
	}

	println(resA)
}
