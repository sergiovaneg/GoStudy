package p0857

import (
	"container/heap"
	"slices"
)

type Worker struct {
	r float64
	q int
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func MincostToHireWorkers(quality []int, wage []int, k int) float64 {
	if k == 1 {
		return float64(slices.Min(wage))
	}

	n := len(quality)
	workers := make([]Worker, n)
	for idx, q := range quality {
		workers[idx] = Worker{r: float64(wage[idx]) / float64(q), q: q}
	}
	slices.SortStableFunc(workers, func(a, b Worker) int {
		if a.r < b.r {
			return -1
		}
		if b.r < a.r {
			return 1
		}
		return 0
	})

	max_heap := make(MaxHeap, k)
	q_sum, r_max := 0, 0.

	for idx := 0; idx < k; idx++ {
		r_max = max(r_max, workers[idx].r)
		q_sum += workers[idx].q
		max_heap[idx] = workers[idx].q
	}
	heap.Init(&max_heap)

	res := r_max * float64(q_sum)

	for idx := k; idx < n; idx++ {
		r_max = max(r_max, workers[idx].r)
		q_sum += workers[idx].q - heap.Pop(&max_heap).(int)
		heap.Push(&max_heap, workers[idx].q)
		res = min(res, r_max*float64(q_sum))
	}

	return res
}
