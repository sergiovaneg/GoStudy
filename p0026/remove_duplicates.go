package p0026

func RemoveDuplicates(nums []int) int {
	// Main loop
	var idx, k int
	for idx, k = 1, len(nums); idx < k; idx++ {
		// Removal optimized for blocks of repeated numbers
		// (worst case: single copy for every number)
		if nums[idx] == nums[idx-1] {
			// comp instantiated to avoid multiple access to slice
			comp, ub := nums[idx], idx+1
			for ub < k {
				if nums[ub] != comp {
					break
				}
				ub++
			}

			if ub == k { // Limit case
				nums = nums[:idx]
			} else {
				nums = append(nums[:idx], nums[ub:]...)
			}
			k -= (ub - idx)
		}
	}

	return k
}
