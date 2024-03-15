package p0238

func ProductExceptSelf(nums []int) []int {
	out := make([]int, len(nums))
	out[0] = 1
	acc := 1
	for idx := 1; idx < len(nums); idx++ {
		acc *= nums[idx-1]
		out[idx] = acc
	}

	acc = 1
	for idx := len(nums) - 2; idx >= 0; idx-- {
		acc *= nums[idx+1]
		out[idx] *= acc
	}

	return out
}
