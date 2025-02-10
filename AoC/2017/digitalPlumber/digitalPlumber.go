package main

import (
	"bufio"
	"os"
	"regexp"
	"slices"
	"strconv"

	"github.com/sergiovaneg/GoStudy/utils"
)

type Graph map[int][]int

func (g *Graph) connect(src, dst int) {
	if _, ok := (*g)[src]; !ok {
		(*g)[src] = make([]int, 0)
	}
	(*g)[src], _ = utils.SortedUniqueInsert((*g)[src], dst)
}

func (g Graph) findConnected(id int) []int {
	component, queue := make([]int, len(g[id])), make([]int, len(g[id]))
	copy(component, g[id])
	component, _ = utils.SortedUniqueInsert(component, id)

	copy(queue, g[id])

	var ok bool
	for i := 0; i < len(queue); i++ {
		for _, cand := range g[queue[i]] {
			component, ok = utils.SortedUniqueInsert(component, cand)
			if !ok {
				queue = append(queue, cand)
			}
		}
	}

	return component
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	g := make(Graph)
	for scanner.Scan() {
		literals := regexp.MustCompile(`(\d+)`).FindAllString(scanner.Text(), -1)
		nums := make([]int, len(literals))
		for i, lit := range literals {
			nums[i], _ = strconv.Atoi(lit)
		}

		for _, num := range nums[1:] {
			g.connect(nums[0], num)
			g.connect(num, nums[0])
		}
	}

	println(len(g.findConnected(0)))

	marked, acc := make([]int, 0), 0
	for id := range g {
		if _, ok := slices.BinarySearch(marked, id); !ok {
			acc++
			marked = utils.SortedMerge(marked, g.findConnected(id))
		}
	}
	println(acc)
}
