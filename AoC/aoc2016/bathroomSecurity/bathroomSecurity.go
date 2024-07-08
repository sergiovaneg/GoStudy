package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func updatePositionA(x0 [2]int, line string) (x [2]int) {
	x[0], x[1] = x0[0], x0[1]

	for _, inst := range line {
		switch inst {
		case 'U':
			if x[0] > 0 {
				x[0]--
			}
		case 'D':
			if x[0] < 2 {
				x[0]++
			}
		case 'L':
			if x[1] > 0 {
				x[1]--
			}
		case 'R':
			if x[1] < 2 {
				x[1]++
			}
		}
	}

	return
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func updatePositionB(x0 [2]int, line string) (x [2]int) {
	x[0], x[1] = x0[0], x0[1]

	for _, inst := range line {
		// Using Manhattan Distance from center
		switch inst {
		case 'U':
			if absInt(x[0]-1)+absInt(x[1]) < 3 {
				x[0]--
			}
		case 'D':
			if absInt(x[0]+1)+absInt(x[1]) < 3 {
				x[0]++
			}
		case 'L':
			if absInt(x[0])+absInt(x[1]-1) < 3 {
				x[1]--
			}
		case 'R':
			if absInt(x[0])+absInt(x[1]+1) < 3 {
				x[1]++
			}
		}
	}

	return
}

func main() {
	keypadA := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	keypadB := [][]int{
		{-1, -1, 1, -1, -1},
		{-1, 2, 3, 4, -1},
		{5, 6, 7, 8, 9},
		{-1, 10, 11, 12, -1},
		{-1, -1, 13, -1, -1},
	}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var codeA, codeB [5]int
	xA, xB := [2]int{1, 1}, [2]int{0, -2}
	for idx := 0; scanner.Scan(); idx++ {
		xA = updatePositionA(xA, scanner.Text())
		xB = updatePositionB(xB, scanner.Text())

		codeA[idx] = keypadA[xA[0]][xA[1]]
		codeB[idx] = keypadB[xB[0]+2][xB[1]+2]
	}

	fmt.Println(codeA)
	fmt.Printf("%x\n", codeB)
}
