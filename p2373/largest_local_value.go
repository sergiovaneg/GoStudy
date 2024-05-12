package p2373

import "slices"

func LargestLocal(grid [][]int) [][]int {
	nm2 := len(grid) - 2
	res := make([][]int, nm2)

	for i := 0; i < nm2; i++ {
		res[i] = make([]int, nm2)
		for j := 0; j < nm2; j++ {
			res[i][j] = max(
				slices.Max(grid[i][j:j+3]),
				slices.Max(grid[i+1][j:j+3]),
				slices.Max(grid[i+2][j:j+3]))
		}
	}

	return res
}
