package p0834

func depthFirstInit(
	graph map[int][]int,
	node, parent int,
	count, res []int) {
	for _, child := range graph[node] {
		if child != parent {
			depthFirstInit(graph, child, node, count, res)
			count[node] += count[child]
			res[node] += res[child] + count[child]
		}
	}
}

func depthFirsUpdate(
	graph map[int][]int,
	node, parent int,
	count, res []int) {
	for _, child := range graph[node] {
		if child != parent {
			res[child] = res[node] - count[child] + (len(res) - count[child])
			depthFirsUpdate(graph, child, node, count, res)
		}
	}
}

func SumOfDistancesInTree(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	graph := make(map[int][]int, n)
	for _, edge := range edges {
		if graph[edge[0]] == nil {
			graph[edge[0]] = make([]int, 1)
			graph[edge[0]][0] = edge[1]
		} else {
			graph[edge[0]] = append(graph[edge[0]], edge[1])
		}
		if graph[edge[1]] == nil {
			graph[edge[1]] = make([]int, 1)
			graph[edge[1]][0] = edge[0]
		} else {
			graph[edge[1]] = append(graph[edge[1]], edge[0])
		}
	}

	// Aux vars
	count, res := make([]int, n), make([]int, n)
	for idx := range count {
		count[idx] = 1
	}

	depthFirstInit(graph, 0, -1, count, res)
	depthFirsUpdate(graph, 0, -1, count, res)

	return res
}
