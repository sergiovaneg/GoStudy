package p1915

func WonderfulSubstrings(word string) int64 {
	count := make([]int64, 1<<10)
	count[0] = 1

	mask := 0x00

	var res int64

	for idx := range word {
		mask ^= 1 << (word[idx] - 'a')
		res += count[mask]
		for shift := 0; shift < 10; shift++ {
			res += count[mask^(1<<shift)]
		}
		count[mask]++
	}

	return res
}
