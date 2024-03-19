package p0621

import "slices"

func LeastInterval(tasks []byte, n int) int {
	task_count_map := make(map[byte]int)
	var unique_tasks []byte
	for _, task_id := range tasks {
		if task_count_map[task_id] == 0 {
			unique_tasks = append(unique_tasks, task_id)
		}
		task_count_map[task_id]++
	}

	cycle_count := 0
	task_idx := 0
	for {
		slices.SortFunc(unique_tasks, func(a, b byte) int {
			return task_count_map[b] - task_count_map[a]
		})
		cd := n + 1
		for _, task_id := range unique_tasks {
			if task_count_map[task_id] > 0 {
				task_count_map[task_id]--
				cycle_count++
				task_idx++
				cd--
			} else {
				continue
			}
			if cd == 0 {
				break
			}
		}
		if task_idx < len(tasks) {
			cycle_count += cd
		} else {
			break
		}
	}
	return cycle_count
}
