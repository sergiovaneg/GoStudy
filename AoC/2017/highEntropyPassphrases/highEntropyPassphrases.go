package main

import (
	"bufio"
	"os"
	"slices"
	"strings"
)

type PreProcessor func(string) string

func isValidPassphrase(line string, f PreProcessor) bool {
	words := strings.Split(line, " ")

	for i, w := range words {
		words[i] = f(w)
	}

	slices.Sort(words)

	for i, w := range words[1:] {
		if w == words[i] {
			return false
		}
	}

	return true
}

func runeSorter(word string) string {
	aux := []rune(word)
	slices.Sort(aux)
	return string(aux)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// n, _ := utils.LineCounter(file)

	var resA, resB int
	for scanner.Scan() {
		if isValidPassphrase(scanner.Text(), func(s string) string { return s }) {
			resA++
			if isValidPassphrase(scanner.Text(), runeSorter) {
				resB++
			}
		}
	}

	println(resA)
	println(resB)
}
