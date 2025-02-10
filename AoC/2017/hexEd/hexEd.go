package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

type Coordinate [3]int

func (z0 Coordinate) step(dir string) (z Coordinate) {
	switch dir {
	case "n":
		z = z0
		z[2]++
	case "s":
		z = z0
		z[2]--
	case "ne":
		z = Coordinate{
			1 - z0[0],
			z0[0] + z0[1],
			z0[0] + z0[2],
		}
	case "nw":
		z = Coordinate{
			1 - z0[0],
			z0[0] + z0[1] - 1,
			z0[0] + z0[2],
		}
	case "se":
		z = Coordinate{
			1 - z0[0],
			z0[0] + z0[1],
			z0[0] + z0[2] - 1,
		}
	case "sw":
		z = Coordinate{
			1 - z0[0],
			z0[0] + z0[1] - 1,
			z0[0] + z0[2] - 1,
		}
	}

	return z
}

func distance(x, y Coordinate) int {
	u := (x[2] - y[2]) - (x[1] - y[1])
	v := (x[0] - y[0]) + (x[1]+y[1])<<1

	if u*v > 0 {
		return utils.AbsInt(u) + utils.AbsInt(v)
	} else {
		return max(utils.AbsInt(u), utils.AbsInt(v))
	}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	var z Coordinate
	var rB int
	for _, dir := range strings.Split(scanner.Text(), ",") {
		z = z.step(dir)
		if aux := distance(z, Coordinate{}); aux > rB {
			rB = aux
		}
	}

	println(distance(z, Coordinate{}))
	println(rB)
}
