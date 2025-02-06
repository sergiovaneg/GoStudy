package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Program struct {
	name     string
	weight   int
	children []*Program
}

func parseTree(scanner *bufio.Scanner) *Program {
	auxMap := make(map[string]*Program)
	var rName string

	for scanner.Scan() {
		pString := strings.Split(
			strings.SplitN(scanner.Text(), " -> ", 1)[0], " ")
		pName := pString[0]
		pWeight, _ := strconv.Atoi(pString[1][1 : len(pString[1])-1])

		pNode, ok := auxMap[pName]
		if !ok {
			pNode = &Program{name: pName}
			auxMap[pName] = pNode
		}
		pNode.weight = pWeight

		cNames := strings.Split(scanner.Text(), " -> ")[1:]
		if len(cNames) == 0 {
			continue
		}

		cNames = strings.Split(cNames[0], ", ")
		pNode.children = make([]*Program, len(cNames))

		for i, cName := range cNames {
			cNode, ok := auxMap[cName]
			if !ok {
				cNode = &Program{name: cName}
				auxMap[cName] = cNode
			}

			pNode.children[i] = cNode
		}

		if rName == "" || slices.Contains(cNames, rName) {
			rName = pName
		}
	}

	return auxMap[rName]
}

func (node *Program) getTotalWeight() int {
	acc := node.weight
	for _, child := range node.children {
		acc += child.getTotalWeight()
	}

	return acc
}

func (node *Program) recursiveCorrect() bool {
	if node.children == nil {
		return false
	}

	for _, child := range node.children {
		if child.recursiveCorrect() {
			return true
		}
	}

	categories := make(map[int][]*Program)

	for _, child := range node.children {
		w := child.getTotalWeight()
		_, ok := categories[w]
		if ok {
			categories[w] = append(categories[w], child)
		} else {
			categories[w] = []*Program{child}
		}
	}

	if len(categories) == 1 {
		return false
	}

	correct, wrong := -1, -1
	for k, v := range categories {
		if len(v) == 1 {
			wrong = k
		} else {
			correct = k
		}
	}

	fmt.Printf("The weight of %v should be %v.\n",
		categories[wrong][0].name,
		correct-wrong+categories[wrong][0].weight)

	return true
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	parent := parseTree(scanner)
	if parent == nil {
		panic("No parent node.")
	}

	println(parent.name)

	parent.recursiveCorrect()
}
