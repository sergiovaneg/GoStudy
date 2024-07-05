package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strconv"

	"github.com/sergiovaneg/GoStudy/utils"
)

const NGroups = 4

type Configuration struct {
	count        int
	entanglement int
}

func cmpConf(a, b Configuration) int {
	if a.count == b.count {
		return a.entanglement - b.entanglement
	}
	return a.count - b.count
}

func getOptimumSetup(weights []int) int {
	target_weight := 0
	for _, weight := range weights {
		target_weight += weight
	}
	target_weight /= NGroups

	candidates := make(map[int]Configuration, 0)

	var dfs func(startIdx, weight, mask, depth, entanglement int)
	dfs = func(startIdx, weight, mask, depth, entanglement int) {
		for idx := startIdx; idx < len(weights); idx++ {
			potential_weight := weight + weights[idx]

			if potential_weight == target_weight {
				candidates[mask|0x01<<idx] = Configuration{
					count:        depth,
					entanglement: entanglement * weights[idx],
				}
				break
			} else if potential_weight < target_weight {
				dfs(idx+1, potential_weight, mask|0x01<<idx,
					depth+1, entanglement*weights[idx])
			} else {
				break
			}
		}
	}
	dfs(0, 0, 0, 1, 1)

	bestConf := &Configuration{
		count:        math.MaxInt,
		entanglement: math.MaxInt,
	}

	var recCheck func(selected []int, level int) bool
	recCheck = func(selected []int, level int) bool {
		for candidate := range candidates {
			compatible := true
			for _, mask := range selected {
				if candidate&mask != 0 {
					compatible = false
					break
				}
			}
			if !compatible {
				continue
			}

			if level == 1 {
				bestConf.count = candidates[selected[0]].count
				bestConf.entanglement = candidates[selected[0]].entanglement
				return true
			} else {
				if recCheck(append(selected, candidate), level-1) {
					return true
				}
			}
		}
		return false
	}
	for mask := range candidates {
		if cmpConf(*bestConf, candidates[mask]) < 0 {
			continue
		}

		selected := make([]int, 1, NGroups)
		selected[0] = mask
		recCheck(selected, NGroups-1)
	}

	return bestConf.entanglement
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	n, _ := utils.LineCounter(file)

	weights := make([]int, 0, n)
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		weights = append(weights, val)
	}

	slices.Sort(weights)
	println(getOptimumSetup(weights))
}
