package main

import (
	"bufio"
	"os"
	"regexp"
	"slices"
	"strconv"

	"github.com/sergiovaneg/GoStudy/utils"
)

type Particle [3][3]int
type Collision struct {
	t   int
	ids [2]int
}

func parseParticle(line string) Particle {
	var z Particle

	for i, match := range regexp.MustCompile(
		`(-{0,1}\d+)`).FindAllString(line, 9) {
		z[i%3][i/3], _ = strconv.Atoi(match)
	}

	return z
}

func (z Particle) manhattanAcc() int {
	var res int
	for _, ax := range z {
		res += utils.AbsInt(ax[2])
	}
	return res
}

func findIntegerRoots(a, b [3]int) []int {
	// x(t) = (a/2)t^2 + (v0+a/2)t + x0
	// -> 2x(t) = at^2 + (2v0+a)t + 2x0
	// Same roots, but avoids truncation
	var p [3]int
	p[0] = (a[0] - b[0]) << 1             // 2c
	p[1] = (a[1]-b[1])<<1 + (a[2] - b[2]) // 2b
	p[2] = a[2] - b[2]                    // 2a

	// Linear case (no acceleration)
	if p[2] == 0 {
		if p[1] == 0 || p[0]%p[1] != 0 {
			return []int{}
		}
		return []int{-p[0] / p[1]}
	}

	discriminant := p[1]*p[1] - p[2]*p[0]<<2

	if discriminant < 0 { // No real roots
		return []int{}
	}

	if discriminant == 0 { // One real root
		return []int{-p[1] / (p[2] << 1)}
	}

	// Two real roots
	dSqrt := utils.ISqrt(discriminant)
	t1 := (-p[1] - dSqrt) / (p[2] << 1)
	t2 := (-p[1] + dSqrt) / (p[2] << 1)

	return []int{t1, t2}
}

func testRoot(a, b [3]int, t int) bool {
	if t < 0 { // Negative time not allowed
		return false
	}

	// Same rationale
	var p [3]int
	p[0] = (a[0] - b[0]) << 1             // 2c
	p[1] = (a[1]-b[1])<<1 + (a[2] - b[2]) // 2b
	p[2] = a[2] - b[2]                    // 2a

	doubleDist := (p[2]*t*t + p[1]*t + p[0])
	return doubleDist == 0
}

func findCollisions(id int, particles []Particle) []Collision {
	collisions := make([]Collision, 0)

	a := particles[id]
	for i, b := range particles[id+1:] {
		times := findIntegerRoots(a[0], b[0])

		for _, t := range times {
			ok := true

			for axIdx := range a {
				if !testRoot(a[axIdx], b[axIdx], t) {
					ok = false
					break
				}
			}

			if ok {
				collisions = append(collisions, Collision{
					t:   t,
					ids: [2]int{id, id + i + 1},
				})
			}
		}
	}

	return collisions
}

func resolveCollisions(particles []Particle) int {
	n := len(particles)
	c := make(chan []Collision, n)

	for i := range n {
		go func(id int) {
			c <- findCollisions(id, particles)
		}(i)
	}

	groupedCollisions := make(map[int][]Collision)
	sortedTimes := make([]int, 0)
	for range particles {
		collisions := <-c
		for _, collision := range collisions {
			groupedCollisions[collision.t] = append(
				groupedCollisions[collision.t], collision)

			sortedTimes, _ = utils.SortedUniqueInsert(
				sortedTimes, collision.t)
		}
	}

	mask := make([]bool, n)
	for _, t := range sortedTimes {
		newlyDestroyed := make([]int, 0)

		for _, collision := range groupedCollisions[t] {
			if mask[collision.ids[0]] || mask[collision.ids[1]] {
				// Already destroyed in earlier collision
				continue
			}

			newlyDestroyed = append(
				newlyDestroyed,
				collision.ids[0],
				collision.ids[1],
			)
		}

		for _, id := range newlyDestroyed {
			mask[id] = true
		}
	}

	remaining := 0
	for _, destroyed := range mask {
		if !destroyed {
			remaining++
		}
	}

	return remaining
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	particles := make([]Particle, 0, n)

	for scanner.Scan() {
		particles = append(
			particles, parseParticle(scanner.Text()))
	}

	minAccParticle := slices.MinFunc(
		particles,
		func(a, b Particle) int {
			return a.manhattanAcc() - b.manhattanAcc()
		})
	println(slices.Index(particles, minAccParticle))

	println(resolveCollisions(particles))
}
