package p0042

import "slices"

func Trap(height []int) int {
	// Definitions
	// N: len(height)
	// k: len(unique(height))
	acc := 0

	// Memory: O(N)
	height = append([]int{slices.Min(height)}, height...)
	leveled_height := make([]int, len(height))
	diff_height := make([]int, len(height)-1)
	sorted_height := make([]int, len(height)-1)
	unique_height := make([]int, 0, len(height)-1)

	// Complexity: O(N log N + N) << O(N log N)
	copy(sorted_height, height[1:])
	slices.Sort(sorted_height)
	unique_height = append(unique_height, sorted_height[0])
	for _, h := range sorted_height[1:] {
		if h != unique_height[len(unique_height)-1] {
			unique_height = append(unique_height, h)
		}
	}

	// Complexity: O(k*N)
	for h_idx := len(unique_height) - 1; h_idx > 0; h_idx-- {
		for idx := 1; idx < len(height); idx++ {
			if height[idx] == unique_height[h_idx] {
				leveled_height[idx] = 1
			}
			diff_height[idx-1] = leveled_height[idx] - leveled_height[idx-1]
		}

		idx_0 := 0
		for idx_0 < len(diff_height)-1 {
			if diff_height[idx_0] == -1 {
				for idx_1, e_1 := range diff_height[idx_0+1:] {
					if e_1 == 1 {
						acc += (idx_1 + 1) * (unique_height[h_idx] - unique_height[h_idx-1])
						idx_0 += idx_1 + 1
						break
					}
				}
			}
			idx_0++
		}
	}

	return acc
}
