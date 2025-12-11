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
	if src == nil || dst == nil {
		return 0
	}

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
	resA := numPaths(g["you"], g["out"], &dp)
	resB := numPaths(
		g["svr"], g["dac"], &dp)*numPaths(
		g["dac"], g["fft"], &dp)*numPaths(
		g["fft"], g["out"], &dp) + numPaths(
		g["svr"], g["fft"], &dp)*numPaths(
		g["fft"], g["dac"], &dp)*numPaths(
		g["dac"], g["out"], &dp)

	println(resA)
	println(resB)
}
