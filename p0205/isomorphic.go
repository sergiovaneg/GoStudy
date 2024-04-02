package p0205

func IsIsomorphic(s string, t string) bool {
	s2t := [128]int{}
	t2s := [128]int{}

	for idx, end := 0, len(s); idx < end; idx++ {
		if s2t[s[idx]] != t2s[t[idx]] {
			return false
		} else {
			s2t[s[idx]], t2s[t[idx]] = idx+1, idx+1
		}
	}
	return true
}
