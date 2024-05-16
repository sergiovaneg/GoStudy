package p0678

func CheckValidString(s string) bool {
	level := 0
	// Try with tiered stacks
	ast_stack := 0

	for _, c := range s {
		switch c {
		case '(':
			level++
		case '*':
			ast_stack++
		case ')':
			if level > 0 {
				level--
			} else if ast_stack > 0 {
				ast_stack--
			} else {
				return false
			}
		}
	}

	for level > 0 && ast_stack > 0 {
		level--
		ast_stack--
	}
	return level == 0
}
