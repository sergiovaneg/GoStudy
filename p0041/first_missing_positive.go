package p0041

func FirstMissingPositive(nums []int) int {
	N := len(nums)

	// Mark out of bounds
	for idx := 0; idx < N; idx++ {
		if (nums[idx] < 1) || (nums[idx] > N) {
			nums[idx] = N + 1
		}
	}

	// Mark numbers within bounds
	for idx := 0; idx < N; idx++ {
		num := nums[idx]
		if num < 0 {
			num = -num
		}

		if num > N {
			continue
		}
		num--
		if nums[num] > 0 {
			nums[num] = -nums[num]
		}
	}

	// Find first missing element within bounds
	for idx := 0; idx < N; idx++ {
		if nums[idx] > 0 {
			return idx + 1
		}
	}

	// Return first element OOB
	return N + 1
}
