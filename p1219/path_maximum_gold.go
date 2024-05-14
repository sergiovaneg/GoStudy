package p1219

import "sync"

type Path [][]bool

func initPath(grid [][]int) Path {
	m, n := len(grid), len(grid[0])
	path := make(Path, m)

	for i, row := range grid {
		path[i] = make([]bool, n)
		for j, val := range row {
			if val > 0 {
				path[i][j] = true
			}
		}
	}

	return path
}

func (path Path) explorePath(grid [][]int, i, j, m, n int) int {
	path[i][j] = false
	var best int

	if i > 0 && path[i-1][j] {
		best = max(best, path.explorePath(grid, i-1, j, m, n))
	}
	if j > 0 && path[i][j-1] {
		best = max(best, path.explorePath(grid, i, j-1, m, n))
	}
	if i < m-1 && path[i+1][j] {
		best = max(best, path.explorePath(grid, i+1, j, m, n))
	}
	if j < n-1 && path[i][j+1] {
		best = max(best, path.explorePath(grid, i, j+1, m, n))
	}

	path[i][j] = true

	return best + grid[i][j]
}

func GetMaximumGold(grid [][]int) int {
	var wg sync.WaitGroup
	c := make(chan int, 25)

	for i, m, n := 0, len(grid), len(grid[0]); i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}

			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				c <- initPath(grid).explorePath(grid, i, j, m, n)
			}(i, j)
		}
	}

	wg.Wait()
	close(c)

	var res int
	for val := range c {
		if val > res {
			res = val
		}
	}

	return res
}
