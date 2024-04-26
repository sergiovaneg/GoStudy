package p1289

func MinFallingPathSum(grid [][]int) int {
	//Early return
	if len(grid) == 1 {
		return grid[0][0]
	}

	get_cost := func(row1, row2 []int) []int {
		cost := make([]int, len(row1))
		for idx1, val1 := range row1 {
			min_value := 100
			for idx2, val2 := range row2 {
				if idx1 == idx2 {
					continue
				}
				if val2 < min_value {
					min_value = val2
				}
			}
			cost[idx1] = val1 + min_value
		}
		return cost
	}

	min_allowed := func(row []int, forbidden int) [2]int {
		min_idx, min_val := -1, 300
		for idx, val := range row {
			if idx == forbidden {
				continue
			}
			if val < min_val {
				min_idx = idx
				min_val = val
			}
		}
		return [2]int{min_idx, min_val}
	}

	// Init
	n, res, forbidden := len(grid), 0, -1

	for i := 0; i < n-2; i++ {
		forbidden = min_allowed(
			get_cost(grid[i], get_cost(grid[i+1], grid[i+2])),
			forbidden)[0]
		res += grid[i][forbidden]
	}

	return res + min_allowed(get_cost(grid[n-2], grid[n-1]), forbidden)[1]
}
