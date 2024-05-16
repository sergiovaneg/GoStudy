package p2812

import (
	"container/heap"
	"math"
)

type Cell struct {
	score int
	loc   [2]int
}

type PriorityQueue []*Cell

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].score != pq[j].score {
		return pq[i].score > pq[j].score
	}
	return pq[i].loc[0]+pq[i].loc[1] > pq[j].loc[0]+pq[j].loc[1]
}
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x any)   { *pq = append(*pq, x.(*Cell)) }

func (pq *PriorityQueue) Pop() any {
	n := len(*pq)
	x := (*pq)[n-1]

	(*pq)[n-1] = nil
	*pq = (*pq)[0 : n-1]

	return x
}

func initThiefs(grid [][]int) [][2]int {
	res := make([][2]int, 0, 1)

	for i, row := range grid {
		for j, cell := range row {
			if cell == 1 {
				res = append(res, [2]int{i, j})
			}
		}
	}

	return res
}

func initPath(n int) [][]bool {
	path := make([][]bool, n)
	for i := range path {
		path[i] = make([]bool, n)
	}

	return path
}

func initScores(grid [][]int, n int) [][]int {
	scores := make([][]int, n)

	for i := 0; i < n; i++ {
		scores[i] = make([]int, n)
		for j := 0; j < n; j++ {
			scores[i][j] = math.MaxInt
		}
	}

	queue := initThiefs(grid)
	for _, thief := range queue {
		scores[thief[0]][thief[1]] = 0
	}

	var next [2]int
	for len(queue) > 0 {
		next, queue = queue[0], queue[1:]
		i, j := next[0], next[1]
		s := scores[i][j]

		if i > 0 && scores[i-1][j] > 1+s {
			scores[i-1][j] = 1 + s
			queue = append(queue, [2]int{i - 1, j})
		}
		if j > 0 && scores[i][j-1] > 1+s {
			scores[i][j-1] = 1 + s
			queue = append(queue, [2]int{i, j - 1})
		}
		if i < n-1 && scores[i+1][j] > 1+s {
			scores[i+1][j] = 1 + s
			queue = append(queue, [2]int{i + 1, j})
		}
		if j < n-1 && scores[i][j+1] > 1+s {
			scores[i][j+1] = 1 + s
			queue = append(queue, [2]int{i, j + 1})
		}
	}

	return scores
}

func MaximumSafenessFactor(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return 0
	}

	path := initPath(n)
	scores := initScores(grid, n)

	pq := make(PriorityQueue, 1)
	pq[0] = &Cell{score: scores[0][0], loc: [2]int{0, 0}}
	heap.Init(&pq)

	var next *Cell
	for pq.Len() > 0 {
		next = heap.Pop(&pq).(*Cell)
		next_loc, next_score := next.loc, next.score

		if next_loc == [2]int{n - 1, n - 1} {
			return next_score
		}
		path[next_loc[0]][next_loc[1]] = true

		if next_loc[0] > 0 && !path[next_loc[0]-1][next_loc[1]] {
			heap.Push(&pq, &Cell{
				score: min(next_score, scores[next_loc[0]-1][next_loc[1]]),
				loc:   [2]int{next_loc[0] - 1, next_loc[1]},
			})
			path[next_loc[0]-1][next_loc[1]] = true
		}
		if next_loc[1] > 0 && !path[next_loc[0]][next_loc[1]-1] {
			heap.Push(&pq, &Cell{
				score: min(next_score, scores[next_loc[0]][next_loc[1]-1]),
				loc:   [2]int{next_loc[0], next_loc[1] - 1},
			})
			path[next_loc[0]][next_loc[1]-1] = true
		}
		if next_loc[0] < n-1 && !path[next_loc[0]+1][next_loc[1]] {
			heap.Push(&pq, &Cell{
				score: min(next_score, scores[next_loc[0]+1][next_loc[1]]),
				loc:   [2]int{next_loc[0] + 1, next_loc[1]},
			})
			path[next_loc[0]+1][next_loc[1]] = true
		}
		if next_loc[1] < n-1 && !path[next_loc[0]][next_loc[1]+1] {
			heap.Push(&pq, &Cell{
				score: min(next_score, scores[next_loc[0]][next_loc[1]+1]),
				loc:   [2]int{next_loc[0], next_loc[1] + 1},
			})
			path[next_loc[0]][next_loc[1]+1] = true
		}
	}

	return -1
}
