package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

type node []*node
type graph map[string]*node

func (g *graph) get(label string) *node {
	if node, ok := (*g)[label]; ok {
		return node
	}

	node := make(node, 0)
	(*g)[label] = &node

	return &node
}

func numPaths(src, dst *node, dp *map[[2]*node]int) int {
	if res, ok := (*dp)[[2]*node{src, dst}]; ok {
		return res
	}

	cnt := 0
	for _, next := range *src {
		if next == dst {
			cnt++
		} else {
			cnt += numPaths(next, dst, dp)
		}
	}

	(*dp)[[2]*node{src, dst}] = cnt
	return cnt
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	g := make(graph, n)

	for scanner.Scan() {
		srcDsts := strings.Split(scanner.Text(), ": ")

		src := g.get(srcDsts[0])
		for dstLabel := range strings.SplitSeq(srcDsts[1], " ") {
			*src = append(*src, g.get(dstLabel))
		}
	}

	dp := make(map[[2]*node]int)
	println(numPaths(g["you"], g["out"], &dp))
}
