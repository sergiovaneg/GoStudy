package p2958

func MaxSubarrayLength(nums []int, k int) int {
	// Early return
	if k == 0 {
		return 0
	}

	max_length := 0
	count := make(map[int]int)
	low, high := 0, 0
	for high < len(nums) {
		// Grow subarray
		last_added := nums[high]
		count[last_added]++
		high++

		// Shrink until valid
		for count[last_added] > k {
			count[nums[low]]--
			low++
		}

		// Update max valid length
		max_length = max(max_length, high-low)
	}

	return max_length
}
