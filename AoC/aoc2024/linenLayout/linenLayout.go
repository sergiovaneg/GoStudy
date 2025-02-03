package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

type Record map[string]int

func (r *Record) isValidPattern(pattern string, options [][]string) int {
	if len(pattern) == 0 {
		return 1
	}
	if val, exists := (*r)[pattern]; exists {
		return val
	}

	count := 0
	for _, flatOptions := range options {
		for _, opt := range flatOptions {
			newPattern, ok := strings.CutPrefix(pattern, opt)

			if !ok {
				continue
			}

			val := r.isValidPattern(newPattern, options)
			if val > 0 {
				count += val
			}

			break
		}

	}

	(*r)[pattern] = count
	return count
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	c := make(chan int, n-2)
	defer close(c)

	scanner.Scan()
	flatOptions := strings.Split(scanner.Text(), ", ")
	options, l, i := make([][]string, 0), -1, -1
	for _, opt := range flatOptions {
		if len(opt) == l {
			options[i] = append(options[i], opt)
		} else {
			i++
			l = len(opt)
			options = append(options, []string{opt})
		}
	}
	slices.SortFunc(flatOptions, func(a, b string) int { return len(b) - len(a) })
	scanner.Scan()

	for scanner.Scan() {
		func(pattern string) {
			r := make(Record)
			c <- r.isValidPattern(pattern, options)
		}(scanner.Text())
	}

	resA, resB := 0, 0
	for range n - 2 {
		if val := <-c; val > 0 {
			resA++
			resB += val
		}
	}

	println(resA)
	println(resB)
}
