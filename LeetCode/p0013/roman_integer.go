package p0013

func RomanToInt(s string) int {
	dict := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := dict[s[len(s)-1]]
	adding := true
	for idx := len(s) - 2; idx >= 0; idx-- {
		if s[idx] != s[idx+1] {
			adding = dict[s[idx]] > dict[s[idx+1]]
		}

		if adding {
			result += dict[s[idx]]
		} else {
			result -= dict[s[idx]]
		}
	}

	return result
}
