package p0005

func isPalindrome(s string) bool {

	l := len(s)
	l_halves := l - l>>1
	for idx1, idx2 := 0, l-1; idx1 < l_halves; idx1, idx2 = idx1+1, idx2-1 {
		if s[idx1] != s[idx2] {
			return false
		}
	}
	return true
}

func LongestPalindrome(s string) string {
	var longest string

	for low, high, l := 0, 1, len(s); high <= l; low, high = low+1, high+1 {
		for {
			if isPalindrome(s[low:high]) {
				longest = s[low:high]
			}

			if low > 0 && high < l && isPalindrome(s[low-1:high+1]) {
				low--
				high++
			} else if low > 0 && isPalindrome(s[low-1:high]) {
				low--
			} else if high < l && isPalindrome(s[low:high+1]) {
				high++
			} else {
				break
			}
		}
	}

	return longest
}
