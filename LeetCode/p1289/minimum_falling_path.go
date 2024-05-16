package p1289

func getCost(row1, row2, cost []int) {
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
}

func minAllowed(row []int, forbidden int) [2]int {
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

func MinFallingPathSum(grid [][]int) int {
	// First early return
	n := len(grid)
	if n == 1 {
		return grid[0][0]
	}

	cost0 := make([]int, n)

	// Second early return
	if n == 2 {
		getCost(grid[0], grid[1], cost0)
		return minAllowed(cost0, -1)[1]
	}

	res, forbidden := 0, -1
	cost1 := make([]int, n)

	for i := 0; i < n-2; i++ {
		// Get the cost of choosing each element on the next row
		getCost(grid[i+1], grid[i+2], cost1)

		// Use the future cost to calculate the current row's cost
		getCost(grid[i], cost1, cost0)

		// Find optimal selection and restrict next selection
		forbidden = minAllowed(cost0, forbidden)[0]

		// Update result
		res += grid[i][forbidden]
	}

	// Skip calculation since already performed
	return res + minAllowed(cost1, forbidden)[1]
}
