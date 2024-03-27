package p0713

func NumSubarrayProductLessThanK(nums []int, k int) int {
	// Early return for infeasible subarrays
	if k <= 1 {
		return 0
	}

	acc := 0
	low, high := 0, 0
	prod := 1
	for high < len(nums) {
		// Increase subarray size
		for high < len(nums) {
			prod *= nums[high]
			high++
			if prod < k {
				acc += (high - low)
			} else {
				break
			}
		}

		// Decrease subarray size if limit exceeded
		if prod >= k {
			for low < high {
				prod /= nums[low]
				low++
				if prod < k {
					break
				}
			}

			acc += (high - low)
		}
	}

	return acc
}
