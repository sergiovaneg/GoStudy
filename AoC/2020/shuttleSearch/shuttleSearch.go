package main

import (
	"bufio"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

func getWaitTime(id, ref int) int {
	q := int(math.Ceil(float64(ref) / float64(id)))

	return q*id - ref
}

func cmpWaitTime(a, b, ref int) int {
	if a == -1 {
		return 1
	}
	if b == -1 {
		return -1
	}

	return getWaitTime(a, ref) - getWaitTime(b, ref)
}

func minimizeTS(buses []int) int {
	// Doesn't work; use the python script
	phaseAcc, periodAcc := 0, 1
	for idx, id := range buses {
		if id == -1 {
			continue
		}
		g, s, _ := utils.ExtGCD(id, periodAcc)
		z := (idx - phaseAcc) / g
		m := z * s

		periodAcc = utils.LCM(periodAcc, id)
		aux := (-m*id + idx)
		phaseAcc = aux % periodAcc
	}

	return phaseAcc
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	ref, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	buses := make([]int, 0)
	for num := range strings.SplitSeq(scanner.Text(), ",") {
		if val, err := strconv.Atoi(num); err == nil {
			buses = append(buses, val)
		} else {
			buses = append(buses, -1)
		}
	}

	optimalId := slices.MinFunc(buses, func(a, b int) int {
		return cmpWaitTime(a, b, ref)
	})

	println(optimalId * getWaitTime(optimalId, ref))
	println(minimizeTS(buses))
}
