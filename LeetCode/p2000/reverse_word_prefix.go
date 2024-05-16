package p2000

import "slices"

func ReversePrefix(word string, ch byte) string {
	s := []rune(word)
	idx := slices.Index(s, rune(ch))

	if idx == -1 {
		return word
	}

	slices.Reverse(s[:idx+1])

	return string(s)
}
