package p2164

import (
	"slices"
)

func SortEvenOdd(nums []int) []int {
	// defaults to 0
	var nums_even []int = make([]int, 0, 50)
	var nums_odd []int = make([]int, 0, 50)

	var current_even bool = true
	var count_even = (len(nums) + 1) >> 1
	var count_odd = len(nums) >> 1

	for _, element := range nums {
		if current_even {
			nums_even = append(nums_even, element)
		} else {
			nums_odd = append(nums_odd, element)
		}
		current_even = !current_even
	}

	slices.Sort(nums_even)
	slices.Sort(nums_odd)

	j := 0
	for idx := 0; idx < count_even; idx++ {
		nums[j] = nums_even[idx]
		j += 2
	}

	j = 1
	for idx := count_odd - 1; idx >= 0; idx-- {
		nums[j] = nums_odd[idx]
		j += 2
	}

	return nums
}
