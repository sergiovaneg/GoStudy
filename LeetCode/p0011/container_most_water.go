package p0011

func MaxArea(height []int) int {
	n := len(height)
	var res, cnd, h0, h1 int

	for idx0, idx1 := 0, n-1; idx0 < idx1; {
		h0, h1 = height[idx0], height[idx1]
		if h0 > h1 {
			cnd = (idx1 - idx0) * h1
			idx1--
		} else {
			cnd = (idx1 - idx0) * h0
			idx0++
		}

		if cnd > res {
			res = cnd
		}
	}

	return res
}
