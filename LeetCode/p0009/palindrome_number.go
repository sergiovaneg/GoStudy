package p0009

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	digits := make([]uint8, 0, 10)
	for x != 0 {
		digits = append(digits, uint8(x%10))
		x /= 10
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}

	return true
}
