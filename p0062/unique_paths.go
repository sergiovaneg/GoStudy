package p0062

type Table map[[2]int]int

func (t *Table) f(x, y int) int {
	if x == 1 || y == 1 {
		return 1
	}

	if x > y {
		x, y = y, x
	}
	if x == 2 {
		return y
	}

	idx := [2]int{x, y}
	if (*t)[idx] == 0 {
		(*t)[idx] = t.f(x-1, y) + t.f(x, y-1)
	}
	return (*t)[idx]
}

func UniquePaths(m int, n int) int {
	if m > n {
		m, n = n, m
	}

	t := make(Table, m*(n+1)-(m*(m+1))>>1-n<<1+1)

	return t.f(m, n)
}
