package p0205

func IsIsomorphic(s string, t string) bool {
	dict_fwd := make(map[byte]struct {
		recorded bool
		val      byte
	})
	dict_bwd := make(map[byte]struct {
		recorded bool
		val      byte
	})

	for idx := 0; idx < len(s); idx++ {
		if !dict_fwd[s[idx]].recorded {
			dict_fwd[s[idx]] = struct {
				recorded bool
				val      byte
			}{true, t[idx]}
		} else if dict_fwd[s[idx]].val != t[idx] {
			return false
		}

		if !dict_bwd[t[idx]].recorded {
			dict_bwd[t[idx]] = struct {
				recorded bool
				val      byte
			}{true, s[idx]}
		} else if dict_bwd[t[idx]].val != s[idx] {
			return false
		}
	}

	return true
}
