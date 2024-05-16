package p2091

import "slices"

func MinimumDeletions(nums []int) int {
	// L := len(nums)
	idx_0 := slices.Index(nums, slices.Min(nums))
	idx_1 := slices.Index(nums, slices.Max(nums))

	if idx_1 < idx_0 { // order
		idx_1 += idx_0
		idx_0 = idx_1 - idx_0
		idx_1 -= idx_0
	}

	// d_0 := idx_0 + 1
	// d_1 := L - idx_1

	if idx_0 == idx_1 {
		// return min(d_0, d_1)
		return min(idx_0+1, len(nums)-idx_1)
	}

	return min(idx_0+1+len(nums)-idx_1, //d_0 + d_1
		1+idx_1,         // d_0 + (idx_1 - idx_0)
		len(nums)-idx_0) // d_1 + (idx_1 - idx_0)
}
