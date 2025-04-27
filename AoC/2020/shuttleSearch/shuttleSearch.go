package main

import (
	"bufio"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/draffensperger/golp"
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
	lp := golp.NewLP(0, len(buses)+1)

	lp.SetInt(0, true)

	for idx, id := range buses {
		if id == -1 {
			lp.AddConstraintSparse(
				[]golp.Entry{{Col: idx + 1, Val: 1.}}, golp.EQ, 0.)
			continue
		}

		lp.AddConstraintSparse(
			[]golp.Entry{{Col: idx + 1, Val: 1.}}, golp.GE, 0.)

		lp.AddConstraintSparse(
			[]golp.Entry{
				{Col: idx + 1, Val: float64(id)},
				{Col: 0, Val: -1.},
			}, golp.EQ, float64(idx))
	}

	obj := make([]float64, len(buses)+1)
	obj[0] = 1.
	lp.SetObjFn(obj)

	lp.Solve()

	return int(math.Round(lp.Variables()[0]))
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
