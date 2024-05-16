package p0089

func stateParse(state []bool) int {
	var res int
	for idx, b := range state {
		if b {
			res += 1 << idx
		}
	}
	return res
}

func stateUpdate(state []bool, counter []int) {
	for idx, b := range state {
		counter[idx]--
		if counter[idx] == 0 {
			counter[idx] = 2 << idx
			state[idx] = !b
		}
	}
}

func GrayCode(n int) []int {
	counter, state := make([]int, n), make([]bool, n)

	for idx := range counter {
		counter[idx] = 1 << idx
	}

	res := make([]int, 1<<n)
	for idx := range res {
		res[idx] = stateParse(state)
		stateUpdate(state, counter)
	}

	return res
}
