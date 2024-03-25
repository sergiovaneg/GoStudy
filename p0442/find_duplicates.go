package p0442

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func FindDuplicates(nums []int) []int {
	result := make([]int, 0, len(nums)/2)
	for idx := 0; idx < len(nums); idx++ {
		aux := absInt(nums[idx])
		if nums[aux-1] < 0 {
			result = append(result, aux)
		}
		nums[aux-1] *= -1
	}
	return result
}
