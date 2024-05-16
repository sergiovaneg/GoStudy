package p1700

func CountStudents(students []int, sandwiches []int) int {
	l := len(sandwiches)
	sw_idx := 0
	st_mark := make([]bool, l)
	for sw_idx < l {
		changed := false
		for st_idx := 0; st_idx < l; st_idx++ {
			if !st_mark[st_idx] && (students[st_idx] == sandwiches[sw_idx]) {
				changed = true
				sw_idx++
				st_mark[st_idx] = true
			}
		}
		if !changed {
			break
		}
	}
	return l - sw_idx
}
