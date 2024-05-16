package p0029

import "strconv"

func Divide(dividend, divisor int) int {
	var negative bool

	// Sign determination
	if dividend >= 0 {
		if divisor < 0 {
			negative = true
			divisor = -divisor
		}
	} else {
		dividend = -dividend
		if divisor > 0 {
			negative = true
		} else {
			divisor = -divisor
		}
	}

	// Init
	res := 0
	dividend_str := strconv.Itoa(dividend)
	sub_str := ""

	// Main loop
	for idx, L := 0, len(dividend_str); idx < L; idx++ {
		aux := res

		// res *= 10
		for i := 0; i < 9; i++ {
			res += aux
		}

		sub_str += dividend_str[idx : idx+1]
		aux, _ = strconv.Atoi(sub_str)

		for aux >= divisor {
			aux -= divisor
			res++
		}
		if aux == 0 {
			sub_str = ""
		} else {
			sub_str = strconv.Itoa(aux)
		}
	}

	if negative {
		return -res
	}
	return res
}
