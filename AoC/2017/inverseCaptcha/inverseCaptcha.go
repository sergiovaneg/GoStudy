package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputString := scanner.Text()

	resA := 0
	for i, r := range inputString[1:] + scanner.Text()[:1] {
		if rune(inputString[i]) == r {
			resA += int(r - '0')
		}
	}
	println(resA)

	resB := 0
	for i, r := range inputString[len(inputString)>>1:] {
		if rune(inputString[i]) == r {
			resB += int(r-'0') << 1
		}
	}
	println(resB)
}
