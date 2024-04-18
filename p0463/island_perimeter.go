package p0463

import "sync"

type GridChecked struct {
	rows, cols   int
	mu_grid      [][]sync.Mutex
	val_grid     [][]int
	mark_grid    [][]bool
	perimeter    int
	mu_perimeter sync.Mutex
}

func (g *GridChecked) exploreCell(row, col int, wg *sync.WaitGroup) {
	defer wg.Done()
	var acc int

	mark_and_spawn := func(row, col int) {
		// Selected is water
		if g.val_grid[row][col] == 0 {
			acc++
			return
		}

		success := g.mu_grid[row][col].TryLock()
		if !success { // Selected already being taken care of
			return
		}

		// Take care of selected
		defer g.mu_grid[row][col].Unlock()
		if !g.mark_grid[row][col] { // Still unchecked
			g.mark_grid[row][col] = true // Mark as inspected
			wg.Add(1)
			go g.exploreCell(row, col, wg)
		}
	}

	// Check top
	if row > 0 {
		mark_and_spawn(row-1, col)
	} else {
		acc++
	}

	// Check Left
	if col > 0 {
		mark_and_spawn(row, col-1)
	} else {
		acc++
	}

	// Check Bottom
	if row < g.rows-1 {
		mark_and_spawn(row+1, col)
	} else {
		acc++
	}

	// Check Right
	if col < g.cols-1 {
		mark_and_spawn(row, col+1)
	} else {
		acc++
	}

	g.mu_perimeter.Lock()
	g.perimeter += acc
	g.mu_perimeter.Unlock()
}

func IslandPerimeter(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	checked := GridChecked{
		rows:      rows,
		cols:      cols,
		val_grid:  grid,
		mark_grid: make([][]bool, rows),
		mu_grid:   make([][]sync.Mutex, rows),
	}
	for row := 0; row < rows; row++ {
		checked.mark_grid[row] = make([]bool, cols)
		checked.mu_grid[row] = make([]sync.Mutex, cols)
	}

	var row, col int
	for row = 0; row < rows; row++ {
		found := false
		for col = 0; col < cols; col++ {
			if grid[row][col] == 1 {
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	var wg sync.WaitGroup
	checked.mark_grid[row][col] = true
	wg.Add(1)
	go checked.exploreCell(row, col, &wg)
	wg.Wait()

	return checked.perimeter
}
