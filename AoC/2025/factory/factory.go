package main

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/draffensperger/golp"
)

type machine struct {
	target   int
	buttons  []int
	joltages []int
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
	joltages = joltages[1 : len(joltages)-1]

	ret.joltages = make([]int, 0)
	for num := range strings.SplitSeq(joltages, ",") {
		val, _ := strconv.Atoi(num)
		ret.joltages = append(ret.joltages, val)
	}

	return ret
}

func (m machine) minPushesA() int {
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

func (m machine) minPushesB() int {
	lp := golp.NewLP(0, len(m.buttons))

	for i, jolt := range m.joltages {
		row := make([]golp.Entry, 0)

		for j, button := range m.buttons {
			if button&(0x01<<i) > 0 {
				row = append(row, golp.Entry{Col: j, Val: 1.})
			}
		}

		lp.AddConstraintSparse(row, golp.EQ, float64(jolt))
	}

	row := make([]float64, len(m.buttons))
	for j := range m.buttons {
		lp.SetInt(j, true)
		row[j] = 1.
	}
	lp.SetObjFn(row)

	lp.Solve()

	return int(math.Round(lp.Objective()))
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
		m := parseMachine(scanner.Text())
		resA += m.minPushesA()
		resB += m.minPushesB()
	}

	println(resA)
	println(resB)
}
