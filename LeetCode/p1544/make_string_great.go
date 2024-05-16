package p1544

func MakeGood(s string) string {
	for idx, L := 1, len(s)-1; idx <= L; {
		var diff byte
		if s[idx] > s[idx-1] {
			diff = s[idx] - s[idx-1]
		} else {
			diff = s[idx-1] - s[idx]
		}
		if diff == 32 {
			if idx == L {
				s = s[:L-1]
				break
			} else {
				s = s[:idx-1] + s[idx+1:]
			}
			L -= 2
			if idx > 1 {
				idx--
			}
		} else {
			idx++
		}
	}
	return s
}
