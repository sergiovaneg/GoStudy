package p1971

func initBooleanPath(n int, edges [][]int, source int) []bool {
	boolean_path := make([]bool, n)
	boolean_path[source] = true

	for _, edge := range edges {
		if edge[0] == source {
			boolean_path[edge[1]] = true
			continue
		}
		if edge[1] == source {
			boolean_path[edge[0]] = true
			continue
		}
	}

	return boolean_path
}

func updateBooleanPath(boolean_path []bool, edges [][]int) bool {
	changed := false

	for _, edge := range edges {
		if !boolean_path[edge[1]] && boolean_path[edge[0]] {
			boolean_path[edge[1]] = true
			changed = true
			continue
		}
		if !boolean_path[edge[0]] && boolean_path[edge[1]] {
			boolean_path[edge[0]] = true
			changed = true
			continue
		}
	}

	return changed
}

func ValidPath(n int, edges [][]int, source int, destination int) bool {
	boolean_path := initBooleanPath(n, edges, source)

	for updateBooleanPath(boolean_path, edges) {
		if boolean_path[destination] {
			return true
		}
	}

	return boolean_path[destination]
}
