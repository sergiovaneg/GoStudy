package p0930

func NumSubarraysWithSum(nums []int, goal int) int {
	return atMost(nums, goal) - atMost(nums, goal-1)
}

func atMost(nums []int, goal int) int {
	var idx_0, idx_1 = 0, 0
	acc, count := 0, 0

	for idx_1 = 0; idx_1 < len(nums); idx_1++ {
		acc += nums[idx_1]
		for (acc > goal) && (idx_0 <= idx_1) {
			acc -= nums[idx_0]
			idx_0++
		}
		count += idx_1 - idx_0 + 1
	}

	return count
}
