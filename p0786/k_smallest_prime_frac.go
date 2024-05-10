package p0781

import (
	"container/heap"
	"math"
)

type Frac struct {
	val     float64
	num_idx int
	den_idx int
}

type FracHeap []*Frac

func (fh FracHeap) Len() int { return len(fh) }

func (fh FracHeap) Less(i, j int) bool { return fh[i].val < fh[j].val }

func (fh FracHeap) Swap(i, j int) { fh[i], fh[j] = fh[j], fh[i] }

func (fh *FracHeap) Push(x any) { *fh = append(*fh, x.(*Frac)) }

func (fh *FracHeap) Pop() any {
	old := *fh
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*fh = old[:n-1]
	return x
}

func KthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	if n == 2 || k == 1 {
		return []int{arr[0], arr[n-1]}
	}

	arr_log := make([]float64, n)
	for idx, num := range arr {
		arr_log[idx] = math.Log(float64(num))
	}

	fh := make(FracHeap, n-1)
	for idx, val := range arr_log[:n-1] {
		fh[idx] = &Frac{val: val - arr_log[n-1], num_idx: idx, den_idx: n - 1}
	}
	heap.Init(&fh)

	for idx := 0; idx < k-1; idx++ {
		aux_frac := heap.Pop(&fh).(*Frac)

		if aux_frac.den_idx-1 > aux_frac.num_idx {
			heap.Push(&fh, &Frac{
				val:     arr_log[aux_frac.num_idx] - arr_log[aux_frac.den_idx-1],
				num_idx: aux_frac.num_idx,
				den_idx: aux_frac.den_idx - 1,
			})
		}
	}

	res := heap.Pop(&fh).(*Frac)
	return []int{arr[res.num_idx], arr[res.den_idx]}
}
