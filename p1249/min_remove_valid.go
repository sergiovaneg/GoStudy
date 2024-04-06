package p1249

func MinRemoveToMakeValid(s string) string {
	s_len := len(s)

	level := 0
	for idx := 0; idx < s_len; {
		c := s[idx]

		if c == '(' {
			level++
		} else if c == ')' {
			if level > 0 {
				level--
			} else {
				if idx == s_len-1 {
					s = s[:idx]
				} else if idx == 0 {
					s = s[1:]
				} else {
					s = s[:idx] + s[idx+1:]
				}
				s_len--
				continue
			}
		}
		idx++
	}

	for idx := s_len - 1; level > 0; idx-- {
		c := s[idx]

		if c == '(' {
			level--
			if idx == s_len-1 {
				s = s[:idx]
			} else if idx == 0 {
				s = s[1:]
			} else {
				s = s[:idx] + s[idx+1:]
			}
		}
	}

	return s
}
