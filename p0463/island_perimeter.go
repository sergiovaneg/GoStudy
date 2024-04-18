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

func IslandPerimeterParallelRecursive(grid [][]int) int {
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

func IslandPerimeterRecursive(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	mark_grid := make([][]bool, rows)
	var perimeter int

	for row := 0; row < rows; row++ {
		mark_grid[row] = make([]bool, cols)
	}

	var row, col int
	for row = 0; row < rows; row++ {
		found := false
		for col = 0; col < cols; col++ {
			if grid[row][col] == 1 {
				found = true
				mark_grid[row][col] = true
				break
			}
		}
		if found {
			break
		}
	}

	var mark_and_spawn func(int, int)
	var explore_cell func(int, int)

	mark_and_spawn = func(row, col int) {
		// Selected is water
		if grid[row][col] == 0 {
			perimeter++
			return
		}

		if !mark_grid[row][col] { // Not yet inspected
			mark_grid[row][col] = true // Mark as inspected
			explore_cell(row, col)
		}
	}

	explore_cell = func(row, col int) {
		// Check top
		if row > 0 {
			mark_and_spawn(row-1, col)
		} else {
			perimeter++
		}

		// Check Left
		if col > 0 {
			mark_and_spawn(row, col-1)
		} else {
			perimeter++
		}

		// Check Bottom
		if row < rows-1 {
			mark_and_spawn(row+1, col)
		} else {
			perimeter++
		}

		// Check Right
		if col < cols-1 {
			mark_and_spawn(row, col+1)
		} else {
			perimeter++
		}
	}

	explore_cell(row, col)

	return perimeter
}

func IslandPerimeterParallelIterative(grid [][]int) int {
	var perimeter int
	rows, cols := len(grid), len(grid[0])

	var wg sync.WaitGroup
	c := make(chan int, rows)

	wg.Add(rows)
	for row := 0; row < rows; row++ {
		go func(row int, c chan<- int, wg *sync.WaitGroup) {
			defer wg.Done()
			var perimeter int
			for col := 0; col < cols; col++ {
				if grid[row][col] == 0 {
					continue
				}

				// Check Top
				if row == 0 || grid[row-1][col] == 0 {
					perimeter++
				}
				// Check Left
				if col == 0 || grid[row][col-1] == 0 {
					perimeter++
				}
				// Check Bottom
				if row == rows-1 || grid[row+1][col] == 0 {
					perimeter++
				}
				// Check Right
				if col == cols-1 || grid[row][col+1] == 0 {
					perimeter++
				}
			}
			c <- perimeter
		}(row, c, &wg)
	}
	wg.Wait()

	for row := 0; row < rows; row++ {
		perimeter += <-c
	}

	return perimeter
}

func IslandPerimeterIterative(grid [][]int) int {
	var perimeter int

	rows, cols := len(grid), len(grid[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 0 {
				continue
			}

			// Check Top
			if row == 0 || grid[row-1][col] == 0 {
				perimeter++
			}
			// Check Left
			if col == 0 || grid[row][col-1] == 0 {
				perimeter++
			}
			// Check Bottom
			if row == rows-1 || grid[row+1][col] == 0 {
				perimeter++
			}
			// Check Right
			if col == cols-1 || grid[row][col+1] == 0 {
				perimeter++
			}
		}
	}

	return perimeter
}
