package p0042

import "slices"

// Complexity: O(N log N + N) << O(N log N)
func getUniqueDescending(height []int) []int {
	sorted_height := make([]int, len(height))
	unique_height := make([]int, len(height))
	unique_idx := 0

	copy(sorted_height, height)
	slices.SortStableFunc(sorted_height, func(a, b int) int { return b - a })
	unique_height[0] = sorted_height[0]
	for _, h := range sorted_height[1:] {
		if h < unique_height[unique_idx] {
			unique_idx++
			unique_height[unique_idx] = h
		}
	}

	return slices.Clip(unique_height[:unique_idx+1])
}

func Trap(height []int) int {
	// Definitions
	// N: len(height)
	// k: len(unique(height))
	acc := 0

	// Memory: O(N)
	height = append([]int{slices.Min(height)}, height...)
	unique_height := getUniqueDescending(height[1:])
	leveled_height := make([]int, len(height))
	diff_height := make([]int, len(height)-1)

	// Complexity: O(k*N)
	for h_idx, uh := range unique_height[:len(unique_height)-1] {
		for idx, h := range height[1:] {
			if h == uh {
				leveled_height[idx+1] = 1
			}
			diff_height[idx] = leveled_height[idx+1] - leveled_height[idx]
		}

		idx_0 := 0
		for idx_0 < len(diff_height)-1 {
			if diff_height[idx_0] == -1 {
				for idx_1, e_1 := range diff_height[idx_0+1:] {
					if e_1 == 1 {
						acc += (idx_1 + 1) * (uh - unique_height[h_idx+1])
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
