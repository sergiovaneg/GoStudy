package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const InitialPos = 50

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// n, _ := utils.LineCounter(file)

	pos, pwdA, pwdB := InitialPos, 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		delta, err := strconv.Atoi(line[1:])

		if err != nil {
			continue
		}

		switch line[0] {
		case 'L':
			if pos == 0 {
				pwdB += delta / 100
			} else {
				pwdB += (delta + (100 - pos)) / 100
			}
			pos -= delta
		case 'R':
			pwdB += (delta + pos) / 100
			pos += delta
		}

		pos %= 100
		if pos < 0 {
			pos = 100 + pos
		}

		if pos == 0 {
			pwdA++
		}

	}

	fmt.Printf("Part A: %v\n", pwdA)
	fmt.Printf("Part B: %v\n", pwdB)
}
