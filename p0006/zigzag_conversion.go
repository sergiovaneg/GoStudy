package p0006

func Convert(s string, numRows int) string {
	// Early return
	if numRows == 1 {
		return s
	}

	n := len(s)

	res, w_idx := make([]byte, n), 0

	shift1, shift2 := (numRows-1)<<1, 0
	for start_idx := 0; start_idx < numRows; start_idx++ {
		r_idx := start_idx
		shift_sel := true
		for r_idx < n {
			if shift_sel {
				if shift1 > 0 {
					res[w_idx] = s[r_idx]
					r_idx += shift1
					w_idx++
				}
			} else {
				if shift2 > 0 {
					res[w_idx] = s[r_idx]
					r_idx += shift2
					w_idx++
				}
			}
			shift_sel = !shift_sel
		}

		shift1 -= 2
		shift2 += 2
	}

	return string(res)
}
