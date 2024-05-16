package p0085

import "sync"

func growAndCalc(matrix [][]byte, i, j, h, w int, just_horz bool) int {
	var valid_i, valid_j bool
	allow_horz := true
	var result_i, result_j int
	if !just_horz && i+h < len(matrix) {
		valid_i = true
		for idx, lim := j, j+w; idx < lim; idx++ {
			if matrix[i+h][idx] == '0' {
				valid_i = false
				break
			}
		}
		if valid_i {
			result_i = growAndCalc(matrix, i, j, h+1, w, false)
		}
	}
	if !valid_i {
		result_i = h * w
	} else if (h * (len(matrix[i]) - j)) < result_i {
		allow_horz = false
	}

	if allow_horz && j+w < len(matrix[i]) {
		valid_j = true
		for idx, lim := i, i+h; idx < lim; idx++ {
			if matrix[idx][j+w] == '0' {
				valid_j = false
				break
			}
		}
		if valid_j {
			result_j = growAndCalc(matrix, i, j, h, w+1, true)
		}
	}
	if !valid_j {
		result_j = h * w
	}

	return max(result_i, result_j)
}

func MaximalRectangle(matrix [][]byte) int {
	current_best := 0

	var wg sync.WaitGroup

	height, width := len(matrix), len(matrix[0])
	c := make(chan int, width)

	for i := 0; i < height; i++ {
		if (height-i)*width < current_best {
			break
		}
		wg.Add(width)
		for j := 0; j < width; j++ {
			if matrix[i][j] == '0' || (height-i)*(width-j) < current_best {
				c <- 0
				wg.Done()
				continue
			}
			i, j := i, j
			go func(c chan int) {
				defer wg.Done()
				c <- growAndCalc(matrix, i, j, 1, 1, false)
			}(c)
		}
		var res int
		for j := 0; j < width; j++ {
			res = <-c
			if res > current_best {
				current_best = res
			}
		}
	}

	return current_best
}
