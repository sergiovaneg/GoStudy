package p0008

const MAX int32 = 0x7FFFFFFF
const MIN int32 = -MAX ^ 1

func MyAtoi(s string) int {
	n := len(s)
	// Early return for empty string
	if n == 0 {
		return 0
	}

	idx := 0

	// Ignore whitespace
	for s[idx] == ' ' {
		idx++
		if idx == n {
			return 0
		}
	}

	// Get sign
	var negative bool
	if s[idx] == '-' {
		negative = true
		idx++
	} else if s[idx] == '+' {
		idx++
	}

	if idx == n {
		return 0
	}

	// Ignore leading zeroes
	for s[idx] == '0' {
		idx++
		if idx == n {
			return 0
		}
	}

	// Start conversion
	var res, limit uint32
	if negative {
		limit = uint32(MAX) + 1
	} else {
		limit = uint32(MAX)
	}
	for _, digit := range s[idx:] {
		if digit > '9' || digit < '0' {
			break
		}
		num := uint32(digit - '0')
		if res > (limit-num)/10 {
			if negative {
				return int(MIN)
			} else {
				return int(MAX)
			}
		}
		res = res*10 + num
	}

	if negative {
		return -int(res)
	} else {
		return int(res)
	}
}
