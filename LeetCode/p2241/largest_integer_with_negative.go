package p2241

import "slices"

func FindMaxK(nums []int) int {
	slices.Sort(nums)
	idx0, idx1 := 0, len(nums)-1

	for idx0 < idx1 {
		if -nums[idx0] == nums[idx1] {
			return nums[idx1]
		}
		if -nums[idx0] > nums[idx1] {
			idx0++
		} else {
			idx1--
		}
	}

	return -1
}
