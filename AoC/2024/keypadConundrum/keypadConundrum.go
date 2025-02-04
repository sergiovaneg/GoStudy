package main

import (
	"bufio"
	"os"
)

type Sequence interface {
	unroll(bool) Nested
}

type Numeric []rune
type Directional rune
type Nested []Sequence

func (src Nested) unroll(return_single bool) Nested {
	unrolled := make(Nested, len(src))

	for i := range src {
		aux := src.unroll(return_single)
		if return_single {
			aux = aux[:1]
		}
		unrolled[i] = aux
	}

	return unrolled
}

func (code Numeric) unroll(return_single bool) Nested {
	x := [2]int{3, 2}

	for _, dst := range code {

	}

}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
	}
}
