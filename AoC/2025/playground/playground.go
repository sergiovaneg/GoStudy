package main

import (
	"bufio"
	"container/heap"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

type coord [3]int
type pair struct {
	idX, idY int
	d2       int
}
type PairHeap []pair

func (h PairHeap) Len() int           { return len(h) }
func (h PairHeap) Less(i, j int) bool { return h[i].d2 < h[j].d2 }
func (h PairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PairHeap) Push(x any)        { *h = append(*h, x.(pair)) }
func (h *PairHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

const connectionNumber = 1000

func parseCoord(line string) coord {
	var x coord

	for i, num := range strings.Split(line, ",") {
		val, _ := strconv.Atoi(num)
		x[i] = val
	}

	return x
}

func d2(x, y coord) int {
	d := 0

	for idx := range 3 {
		d += (x[idx] - y[idx]) * (x[idx] - y[idx])
	}

	return d
}

func getDistanceHeap(arr []coord) PairHeap {
	n := len(arr)
	h := make(PairHeap, 0, (n*(n<<1))>>1)

	heap.Init(&h)
	for i, x := range arr {
		for j, y := range arr[i+1:] {
			item := pair{idX: i, idY: i + j + 1, d2: d2(x, y)}
			heap.Push(&h, item)
		}
	}

	return h
}

func heapIteration(h *PairHeap, membership *[][]int) pair {
	p := heap.Pop(h).(pair)

	// Locate pre-assigned boxes
	setX, setY := -1, -1
	for id, members := range *membership {
		if _, found := slices.BinarySearch(members, p.idX); found {
			setX = id
		}
		if _, found := slices.BinarySearch(members, p.idY); found {
			setY = id
		}
	}

	if setX > -1 && setY > -1 {
		if setX != setY {
			(*membership)[setX] = utils.SortedMerge(
				(*membership)[setX], (*membership)[setY])
			*membership = slices.Delete(*membership, setY, setY+1)
		}
	} else if setX > -1 {
		(*membership)[setX], _ = utils.SortedUniqueInsert(
			(*membership)[setX], p.idY)
	} else if setY > -1 {
		(*membership)[setY], _ = utils.SortedUniqueInsert(
			(*membership)[setY], p.idX)
	} else {
		*membership = append(*membership, []int{p.idX, p.idY})
	}

	return p
}

func heapConnect(arr []coord) (int, int) {
	h := getDistanceHeap(arr)
	membership := make([][]int, 0)
	var p pair

	for range connectionNumber {
		p = heapIteration(&h, &membership)
	}

	slices.SortFunc(membership, func(a, b []int) int {
		return len(b) - len(a)
	})
	resA := len(membership[0]) * len(membership[1]) * len(membership[2])

	for len(membership[0]) < len(arr) {
		p = heapIteration(&h, &membership)
	}

	return resA, arr[p.idX][0] * arr[p.idY][0]
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	arr := make([]coord, 0, n)

	for scanner.Scan() {
		arr = append(arr, parseCoord(scanner.Text()))
	}

	println(heapConnect(arr))
}
