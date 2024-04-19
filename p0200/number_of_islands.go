package p0200

import (
	"sync"
)

type ParallelGrid struct {
	rows, cols   int
	mu_grid      [][]sync.RWMutex
	grid         [][]byte
	id_grid      [][]int
	interrupt    []bool
	mu_interrupt []sync.RWMutex
}

func (g *ParallelGrid) explore_cell(row, col, id int) bool {
	mu_interrupt := &g.mu_interrupt[id-1]

	mu_interrupt.RLock()
	interrupt := g.interrupt[id-1]
	mu_interrupt.RUnlock()

	if interrupt {
		return false
	}

	mark_and_spawn := func(row, col int) bool {

		if g.grid[row][col] == '0' {
			return true
		}

		mu_grid := &g.mu_grid[row][col]

		mu_grid.RLock()
		observed_id := g.id_grid[row][col]
		mu_grid.RUnlock()

		if observed_id == 0 { // I have to do the job of marking
			mu_grid.Lock()
			g.id_grid[row][col] = id
			mu_grid.Unlock()

			return g.explore_cell(row, col, id)
		}

		if observed_id > id {
			mu_grid.Lock()
			g.id_grid[row][col] = id
			mu_grid.Unlock()

			g.mu_interrupt[observed_id-1].Lock()
			g.interrupt[observed_id-1] = true
			g.mu_interrupt[observed_id-1].Unlock()

			return g.explore_cell(row, col, id)
		}
		if observed_id == id { // Tracing-back
			return true
		}

		// Someone was here before me
		mu_interrupt.Lock()
		g.interrupt[id-1] = true
		mu_interrupt.Unlock()

		return false
	}

	// Check top
	if row > 0 && !mark_and_spawn(row-1, col) {
		return false
	}

	// Check Left
	if col > 0 && !mark_and_spawn(row, col-1) {
		return false
	}

	// Check Bottom
	if row < g.rows-1 && !mark_and_spawn(row+1, col) {
		return false
	}

	// Check Right
	if col < g.cols-1 && !mark_and_spawn(row, col+1) {
		return false
	}

	return true
}

func NumIslands(grid [][]byte) int {
	g := ParallelGrid{
		rows:         len(grid),
		cols:         len(grid[0]),
		mu_grid:      make([][]sync.RWMutex, len(grid)),
		grid:         grid,
		id_grid:      make([][]int, len(grid)),
		interrupt:    make([]bool, len(grid)*len(grid[0])),
		mu_interrupt: make([]sync.RWMutex, len(grid)*len(grid[0])),
	}
	for row := 0; row < g.rows; row++ {
		g.mu_grid[row] = make([]sync.RWMutex, g.cols)
		g.id_grid[row] = make([]int, g.cols)
	}

	id := 1
	c := make(chan bool, g.rows*g.cols) // Oversized buffer

	var wg sync.WaitGroup

	starter := func(row, col int) {
		mu := &g.mu_grid[row][col]

		mu.RLock()
		observed_id := g.id_grid[row][col]
		mu.RUnlock()

		if g.grid[row][col] == '0' || observed_id > 0 {
			return
		}

		wg.Add(1)

		mu.Lock()
		g.id_grid[row][col] = id
		mu.Unlock()

		go func(id int, c chan<- bool) {
			defer wg.Done()
			if g.explore_cell(row, col, id) {
				c <- true
			}
		}(id, c)

		id++
	}

	for row := 0; row < g.rows-g.rows/2; row++ {
		for col := 0; col < g.cols-g.cols/2; col++ {
			starter(row, col)
			starter(g.rows-row-1, col)
			starter(row, g.cols-col-1)
			starter(g.rows-row-1, g.cols-col-1)
		}
	}
	wg.Wait()

	return len(c)
}
