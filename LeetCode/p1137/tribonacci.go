package p1137

func Tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n < 3 {
		return 1
	}
	nums := [3]int{0, 1, 1}
	for m, idx := 3, 0; m < n; m++ {
		nums[idx] = nums[0] + nums[1] + nums[2]
		idx = (idx + 1) % 3
	}

	return nums[0] + nums[1] + nums[2]
}
