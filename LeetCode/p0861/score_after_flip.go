package p0861

func MatrixScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	flipped := make([]bool, m)

	for i := range grid {
		flipped[i] = grid[i][0] == 0
	}

	res := m * (1 << (n - 1))
	var cnt int

	for j := 1; j < n; j++ {
		cnt = 0
		for i := 0; i < m; i++ {
			if flipped[i] {
				if grid[i][j] == 0 {
					cnt++
				}
			} else {
				if grid[i][j] == 1 {
					cnt++
				}
			}
		}
		res += max(cnt, m-cnt) * (1 << (n - j - 1))
	}

	return res
}
