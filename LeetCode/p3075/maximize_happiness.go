package p3075

import "slices"

func MaximumHappinessSum(happiness []int, k int) int64 {
	var res int

	slices.Sort(happiness)
	slices.Reverse(happiness)

	for idx := 0; idx < k; idx++ {
		if aux := happiness[idx] - idx; aux > 0 {
			res += aux
		} else {
			break
		}
	}

	return int64(res)
}
