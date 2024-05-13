package p0050

func MyPow(x float64, n int) float64 {
	if n == 0 {
		return 1.
	}

	if n < 0 {
		x = 1. / x
		n = -n
	}

	res := 1.

	for n > 0 {
		if n%2 == 1 {
			res *= x
			n--
		} else {
			x *= x
			n >>= 1
		}
	}

	return res
}
