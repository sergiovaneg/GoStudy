package p0452

import (
	"slices"
)

func FindMinArrowShots(points [][]int) int {
	if len(points) == 1 {
		return 1
	}

	slices.SortFunc(points, func(a, b []int) int {
		if a[0] == b[0] {
			return a[1] - b[1]
		} else {
			return a[0] - b[0]
		}
	})

	arrow_count := 1
	ub := points[0][1]

	for idx := 0; idx < len(points); idx++ {
		if points[idx][0] > ub {
			arrow_count++
			ub = points[idx][1]
		} else {
			ub = min(ub, points[idx][1])
		}
	}

	return arrow_count
}
