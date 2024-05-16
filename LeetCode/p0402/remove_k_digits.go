package p0402

func removeZeroPadding(num string) string {
	for idx, c := range num {
		if c != '0' {
			return num[idx:]
		}
	}
	return "0"
}

func RemoveKdigits(num string, k int) string {
	L := len(num)
	if L == k {
		return "0"
	}

	for idx := 1; idx < L && k > 0; idx++ {
		ll := idx
		for ll > 0 && num[idx] < num[ll-1] && k > 0 {
			k--
			L--
			ll--
		}
		if ll != idx {
			num = num[:ll] + num[idx:]
			idx = ll
		}
	}
	if k > 0 {
		num = num[:L-k]
	}
	return removeZeroPadding(num)
}
