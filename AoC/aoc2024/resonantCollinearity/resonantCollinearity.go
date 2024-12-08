package main

import (
	"bufio"
	"log"
	"os"

	"github.com/sergiovaneg/GoStudy/utils"
)

type Coord [2]int

func isValidCoord(x Coord, m, n int) bool {
	if x[0] < 0 || x[1] < 0 {
		return false
	}

	if x[0] >= m || x[1] >= n {
		return false
	}

	return true
}

func getUniqueRecord(coords []Coord, m, n, lb, ub int) map[Coord]bool {
	l := len(coords)

	record := make(map[Coord]bool, 2*l*(l-1))
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			d := Coord{
				coords[j][0] - coords[i][0],
				coords[j][1] - coords[i][1],
			}

			for k := lb; k != ub; k++ {
				aux := Coord{
					coords[i][0] - k*d[0],
					coords[i][1] - k*d[1],
				}
				if isValidCoord(aux, m, n) {
					record[aux] = true
				} else {
					break
				}
			}

			for k := lb; k != ub; k++ {
				aux := Coord{
					coords[j][0] + k*d[0],
					coords[j][1] + k*d[1],
				}
				if isValidCoord(aux, m, n) {
					record[aux] = true
				} else {
					break
				}
			}
		}
	}

	return record
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	m, _ := utils.LineCounter(file)

	antennaeRecord, n := make(map[rune][]Coord), -1
	for i := 0; scanner.Scan(); i++ {
		n = len(scanner.Text())
		for j, c := range scanner.Text() {
			if c == '.' {
				continue
			}

			if antennaeRecord[c] == nil {
				antennaeRecord[c] = []Coord{{i, j}}
			} else {
				antennaeRecord[c] = append(antennaeRecord[c], Coord{i, j})
			}
		}
	}

	cA := make(chan map[Coord]bool, len(antennaeRecord))
	cB := make(chan map[Coord]bool, len(antennaeRecord))
	for _, coords := range antennaeRecord {
		go func(coords []Coord) {
			cA <- getUniqueRecord(coords, m, n, 1, 2)
			cB <- getUniqueRecord(coords, m, n, 0, -1)
		}(coords)
	}

	uRA, uRB := make(map[Coord]bool), make(map[Coord]bool)
	for range antennaeRecord {
		for coord := range <-cA {
			uRA[coord] = true
		}
		for coord := range <-cB {
			uRB[coord] = true
		}
	}

	println(len(uRA))
	println(len(uRB))
}
