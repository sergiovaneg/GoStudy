package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func parseSides(line string) (sides [3]int) {
	for idx, num := range regexp.MustCompile("([0-9]+)").FindAllString(line, 3) {
		val, _ := strconv.Atoi(num)
		sides[idx] = val
	}

	return
}

func isValidTriangle(sides [3]int) bool {
	x := sides[0] + sides[1] - sides[2]
	y := sides[0] - sides[1] + sides[2]

	return x > 0 && y > 0 && sides[2]<<1 > y
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	counterA, counterB, buffCounter := 0, 0, 0
	var buffer [3][3]int
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		triangle := parseSides(scanner.Text())
		for idx := range 3 {
			buffer[idx][buffCounter] = triangle[idx]
		}

		if isValidTriangle(triangle) {
			counterA++
		}

		if buffCounter == 2 {
			for idx := range 3 {
				if isValidTriangle(buffer[idx]) {
					counterB++
				}
			}
			buffCounter = 0
		} else {
			buffCounter++
		}
	}

	println(counterA)
	println(counterB)
}
