package p0881

import "slices"

func NumRescueBoats(people []int, limit int) int {
	var res int
	slices.SortFunc(people, func(a, b int) int { return b - a })

	for idx0, idx1 := 0, len(people)-1; idx0 <= idx1; idx0++ {
		res++
		if people[idx1] <= limit-people[idx0] {
			idx1--
		}
	}

	return res
}
