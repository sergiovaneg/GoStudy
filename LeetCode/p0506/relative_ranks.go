package p0506

import (
	"fmt"
	"slices"
)

func FindRelativeRanks(score []int) []string {
	n := len(score)
	indices := make([]int, 0, n)
	for idx := range score {
		indices = append(indices, idx)
	}
	slices.SortStableFunc(
		indices,
		func(a, b int) int { return score[b] - score[a] })

	res := make([]string, n)
	res[indices[0]] = "Gold Medal"
	if n > 1 {
		res[indices[1]] = "Silver Medal"
	}
	if n > 2 {
		res[indices[2]] = "Bronze Medal"
	}
	for i := 3; i < n; i++ {
		res[indices[i]] = fmt.Sprint(i + 1)
	}

	return res
}
