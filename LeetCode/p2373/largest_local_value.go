package p2373

func LargestLocal(grid [][]int) [][]int {
	nm2 := len(grid) - 2
	res := make([][]int, nm2)

	for i := 0; i < nm2; i++ {
		res[i] = make([]int, nm2)
		for j := 0; j < nm2; j++ {
			res[i][j] = max(
				grid[i][j], grid[i][j+1], grid[i][j+2],
				grid[i+1][j], grid[i+1][j+1], grid[i+1][j+2],
				grid[i+2][j], grid[i+2][j+1], grid[i+2][j+2])
		}
	}

	return res
}
