package p1992

import "slices"

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

func FindFarmland(land [][]int) [][]int {
	m, n := len(land), len(land[0])
	area := m * n
	res := make([][]int, 0, area-area/2)

	skips := make([][3]int, 0, n)
	skips = append(skips, [3]int{m, 0, 0}, [3]int{m, n, n})

	for i := 0; i < m; i++ {
		for skip_idx := 1; skip_idx < len(skips); skip_idx++ {
			for j := skips[skip_idx-1][2]; j < skips[skip_idx][1]; j++ {
				if land[i][j] == 1 {
					br := findWatson(land, i, j, m, n)
					res = append(res, []int{i, j, br[0], br[1]})
					skips = slices.Insert(skips, skip_idx, [3]int{br[0], j, br[1] + 1})
					skip_idx--
				}
			}
		}
		for skip_idx := 1; skip_idx < len(skips)-1; skip_idx++ {
			if i == skips[skip_idx][0] {
				skips = slices.Delete(skips, skip_idx, skip_idx+1)
				skip_idx--
			}
		}
	}

	return res
}
