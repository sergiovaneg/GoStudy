package p0058

func LengthOfLastWord(s string) int {
	ul := len(s) - 1
	for ul >= 0 && s[ul] == ' ' {
		ul--
	}
	ll := ul - 1
	for ll >= 0 && s[ll] != ' ' {
		ll--
	}
	return ul - ll
}
