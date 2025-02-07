package main

import (
	"bufio"
	"os"
	"regexp"
	"slices"
)

func cleanStream(stream string) (string, int) {
	stream = regexp.MustCompile(`!.{1}`).ReplaceAllString(stream, "")

	matches := regexp.MustCompile(`<[^>]*>`).FindAllStringIndex(stream, -1)
	slices.Reverse(matches)

	var size int
	for _, match := range matches {
		size += match[1] - match[0] - 2
		stream = stream[:match[0]] + stream[match[1]:]
	}

	return stream, size
}

func scoreStream(stream string) int {
	var score, level int

	for _, c := range stream {
		if c == '{' {
			level++
			score += level
		} else if c == '}' {
			level--
		}
	}

	return score
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	stream, garbageSize := cleanStream(scanner.Text())
	println(scoreStream(stream))
	println(garbageSize)
}
