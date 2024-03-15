package p0238

func ProductExceptSelf(nums []int) []int {
	// prefix
	out := make([]int, len(nums))
	out[0] = nums[0]
	for idx, num := range nums[1:] {
		out[idx+1] = out[idx] * num
	}

	// nums -> suffix
	for idx := len(nums) - 2; idx > 0; idx-- {
		nums[idx] *= nums[idx+1]
	}

	// prefix -> out
	out[len(out)-1] = out[len(out)-2]
	for idx := len(out) - 2; idx > 0; idx-- {
		out[idx] = out[idx-1] * nums[idx+1]
	}
	out[0] = nums[1]
	return out
}
