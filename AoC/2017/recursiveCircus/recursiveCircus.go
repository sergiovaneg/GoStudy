package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Program struct {
	name     string
	weight   int
	parent   *Program
	children []*Program
}

func parseTree(scanner *bufio.Scanner) *Program {
	auxMap := make(map[string]*Program)

	for scanner.Scan() {
		parentString := strings.Split(
			strings.SplitN(scanner.Text(), " -> ", 1)[0], " ")
		parentName := parentString[0]
		parentWeight, _ := strconv.Atoi(parentString[1][1 : len(parentString[1])-1])

		parentNode, ok := auxMap[parentName]
		if !ok {
			parentNode = &Program{name: parentName}
			auxMap[parentName] = parentNode
		}
		parentNode.weight = parentWeight

		childrenNames := strings.Split(scanner.Text(), " -> ")[1:]
		if len(childrenNames) == 0 {
			continue
		}

		childrenNames = strings.Split(childrenNames[0], ", ")
		parentNode.children = make([]*Program, len(childrenNames))

		for i, childName := range childrenNames {
			childNode, ok := auxMap[childName]
			if !ok {
				childNode = &Program{name: childName}
				auxMap[childName] = childNode
			}

			parentNode.children[i] = childNode
			childNode.parent = parentNode
		}
	}

	for _, node := range auxMap {
		if node.parent == nil {
			return node
		}
	}

	return nil
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
