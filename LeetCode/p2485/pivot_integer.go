package p2485

func PivotInteger(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return -1
	}

	n_halves := n >> 1
	n_minus_n_halves := n - n_halves

	cumsum := make([]int, n_minus_n_halves)
	cumsum[0] = 1

	var idx int

	for idx = 2; idx < n_halves+2; idx++ {
		cumsum[0] += idx
	}

	for idx = 1; idx < n_minus_n_halves; idx++ {
		cumsum[idx] = cumsum[idx-1] + (idx + n_halves + 1)
	}

	idx = n_minus_n_halves - 2
	thr := cumsum[n_minus_n_halves-1] + n_halves + 1
	for (cumsum[idx] + cumsum[idx] - idx) > thr {
		idx--
	}
	if (cumsum[idx] + cumsum[idx] - idx) == thr {
		return idx + n_halves + 1
	}
	return -1
}
