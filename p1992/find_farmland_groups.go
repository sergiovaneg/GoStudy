package p1992

import (
	"slices"
	"sync"
)

func findWatson(land [][]int, i, j, m, n int) [2]int {
	br := [2]int{i, j}
	moving_down := !((i == m-1) || (land[i+1][j] == 0))
	moving_right := !((j == n-1) || (land[i][j+1] == 0))
	for moving_right || moving_down {
		if moving_down {
			br[0]++
			moving_down = !((br[0] == m-1) || (land[br[0]+1][j] == 0))
		}
		if moving_right {
			br[1]++
			moving_right = !((br[1] == n-1) || (land[i][br[1]+1] == 0))
		}
	}

	return br
}

func exploreRectangle(land [][]int, m, n,
	i_low, j_low, i_high, j_high int) [][]int {
	area := (i_high - i_low) * (n - j_low)
	res := make([][]int, 0, area-area>>1)

	var wg sync.WaitGroup
	c := make(chan [][]int, 2*(i_high-i_low))
	spawn := func(c chan<- [][]int, i_low, j_low, i_high, j_high int) {
		defer wg.Done()
		c <- exploreRectangle(land, m, n, i_low, j_low, i_high, j_high)
	}

	for i := i_low; i < i_high; i++ {
		j := slices.Index(land[i][j_low:j_high], 1) + j_low

		if j != j_low-1 && (i == 0 || land[i-1][j] == 0) {
			// Get rectangle
			br := findWatson(land, i, j, m, n)
			res = append(res, []int{i, j, br[0], br[1]})

			// Spawn to the right
			if br[1] < n-1 {
				wg.Add(1)
				go spawn(c, i, br[1]+1, br[0]+1, j_high)
			}

			// Spawn to the left
			if j > 0 && br[0] != i {
				wg.Add(1)
				go spawn(c, i+1, j_low, br[0]+1, j)
			}

			// Skip to the bottom
			i = br[0]
		}
	}

	wg.Wait()
	for len(c) > 0 {
		res = append(res, <-c...)
	}

	return res
}

func FindFarmland(land [][]int) [][]int {
	m, n := len(land), len(land[0])

	return exploreRectangle(land, m, n, 0, 0, m, n)
}
