package p0004

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)

	if n1 > n2 {
		return FindMedianSortedArrays(nums2, nums1)
	}
	if n1 == 0 { // nums1 is always the smallest
		if n2%2 == 1 {
			return float64(nums2[n2>>1])
		} else {
			return float64(nums2[n2>>1]+nums2[n2>>1-1]) / 2.
		}
	}

	n := n1 + n2
	total_left_size := (n + 1) >> 1

	low_idx, high_idx := 0, n1

	for low_idx <= high_idx {
		mid_idx_1 := (low_idx + high_idx) >> 1
		mid_idx_2 := total_left_size - mid_idx_1

		var l1, l2, r1, r2 int = -1e6, -1e6, 1e6, 1e6

		if mid_idx_1 < n1 {
			r1 = nums1[mid_idx_1]
		}
		if mid_idx_2 < n2 {
			r2 = nums2[mid_idx_2]
		}
		if mid_idx_1 > 0 {
			l1 = nums1[mid_idx_1-1]
		}
		if mid_idx_2 > 0 {
			l2 = nums2[mid_idx_2-1]
		}

		if l1 <= r2 && l2 <= r1 { // Found correct partition
			if n%2 == 1 {
				return float64(max(l1, l2))
			} else {
				return float64(max(l1, l2)+min(r1, r2)) / 2.
			}
		} else if l1 > r2 { // Move towards the left side of nums1
			high_idx = mid_idx_1 - 1
		} else { // Move towards the right side of nums1
			low_idx = mid_idx_1 + 1
		}
	}

	return 0. // The arays were not ordered
}
