package main

import (
	"bufio"
	"math/bits"
	"os"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

type Component [2]int
type Bridge uint //Bitmask
func (b0 Bridge) contains(b1 Bridge) bool {
	return (b0 | b1) == b0
}

type State struct {
	mem       map[int][][2]Bridge
	inventory []Component
}
type bridgeCmpFunc func(State, Bridge, Bridge) int

func (dp State) genFullMask() Bridge {
	var mask Bridge

	for idx := range dp.inventory {
		mask |= 0x01 << idx
	}

	return mask
}

func (dp *State) getStrength(b Bridge) int {
	s := 0

	for idx, segment := range dp.inventory {
		if (0x01<<idx)&b != 0 {
			s += segment[0] + segment[1]
		}
	}

	return s
}

func (dp *State) optimize(port int, avail Bridge, cmp bridgeCmpFunc) Bridge {
	for _, prevCalls := range dp.mem[port] {
		if !avail.contains(prevCalls[1]) {
			continue
		}

		if !prevCalls[0].contains(avail) {
			continue
		}

		return prevCalls[1]
	}

	var bestBridge Bridge

	for idx, segment := range dp.inventory {
		if (0x01<<idx)&avail == 0 { // Piece not available
			continue
		}

		var next int
		if segment[0] == port {
			next = segment[1]
		} else if segment[1] == port {
			next = segment[0]
		} else { // Piece not compatible
			continue
		}

		bridge := dp.optimize(next, avail&^(0x01<<idx), cmp) | (0x01 << idx)
		if cmp(*dp, bridge, bestBridge) > 0 {
			bestBridge = bridge
		}
	}

	if _, ok := dp.mem[port]; !ok {
		dp.mem[port] = make([][2]Bridge, 0)
	}
	dp.mem[port] = append(dp.mem[port], [2]Bridge{avail, bestBridge})

	return bestBridge
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	dp := State{
		mem:       make(map[int][][2]Bridge),
		inventory: make([]Component, 0, n),
	}

	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), "/")
		s0, _ := strconv.Atoi(nums[0])
		s1, _ := strconv.Atoi(nums[1])
		dp.inventory = append(dp.inventory, Component{s0, s1})
	}

	bridge := dp.optimize(0, dp.genFullMask(),
		func(s State, b1, b2 Bridge) int {
			return s.getStrength(b1) - s.getStrength(b2)
		})
	println(dp.getStrength(bridge))

	dp.mem = make(map[int][][2]Bridge) // Clear memory
	bridge = dp.optimize(0, dp.genFullMask(),
		func(s State, b1, b2 Bridge) int {
			l1, l2 := bits.OnesCount(uint(b1)), bits.OnesCount(uint(b2))
			if l1 == l2 {
				return s.getStrength(b1) - s.getStrength(b2)
			}
			return l1 - l2
		})
	println(dp.getStrength(bridge))
}
