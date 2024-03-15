package p0930

func NumSubarraysWithSum(nums []int, goal int) int {
	acc := 0
	cnt := 0
	for idx_0 := 0; idx_0 < len(nums); idx_0++ {
		acc = nums[idx_0]
		if acc == goal {
			cnt++
		} else if acc > goal {
			continue
		}
		for idx_1 := idx_0 + 1; idx_1 < len(nums); idx_1++ {
			acc += nums[idx_1]
			if acc == goal {
				cnt++
			} else if acc > goal {
				break
			}
		}
	}
	return cnt
}
