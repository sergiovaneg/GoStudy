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

func (z Particle) manhattanAcceleration() int {
	var res int
	for _, ax := range z {
		res += utils.AbsInt(ax[2])
	}
	return res
}

func findIntegerRoots(a, b [3]int) []int {
	var delta [3]int

	for i := range 3 {
		delta[i] = a[i] - b[i]
	}

	if delta[2] == 0 {
		if delta[1] == 0 || delta[0]%delta[1] != 0 {
			return []int{}
		}
		return []int{-delta[0] / delta[1]}
	}

	discriminant := delta[1]*delta[1] - delta[2]*delta[0]<<1

	if discriminant < 0 {
		return []int{}
	}

	if discriminant == 0 {
		return []int{-delta[1] / delta[2]}
	}

	dSqrt := utils.ISqrt(discriminant)
	t1 := (-delta[1] - dSqrt) / delta[2]
	t2 := (-delta[1] + dSqrt) / delta[2]

	d1 := delta[2]*t1*t1>>1 + delta[1]*t1 + delta[0]
	d2 := delta[2]*t2*t2>>1 + delta[1]*t2 + delta[0]
	println(d1)
	println(d2)

	return []int{t1, t2}
}

func testRoot(a, b [3]int, t int) bool {
	if t < 0 {
		return false
	}

	var d [3]int

	for i := range 3 {
		d[i] = a[i] - b[i]
	}

	distance := ((d[2]*t*t)>>1 + d[1]*t + d[0])
	return distance == 0
}

func findCollisions(id int, particles []Particle) []Collision {
	collisions := make([]Collision, 0)

	a := particles[id]
	for i := id + 1; i < len(particles); i++ {
		b := particles[i]
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
					ids: [2]int{id, i},
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
		func(id int) {
			c <- findCollisions(id, particles)
		}(i)
	}

	groupedCollisions := make(map[int][]Collision)
	for range particles {
		collisions := <-c
		for _, collision := range collisions {
			groupedCollisions[collision.t] = append(
				groupedCollisions[collision.t], collision)
		}
	}

	sortedTimes := make([]int, 0)
	for k := range groupedCollisions {
		sortedTimes = append(sortedTimes, k)
	}
	slices.Sort(sortedTimes)

	mask := make([]bool, n)
	for _, t := range sortedTimes {
		newlyDestroyed := make([]int, 0)
		for _, collision := range groupedCollisions[t] {
			if mask[collision.ids[0]] || mask[collision.ids[1]] {
				continue
			}

			newlyDestroyed = append(
				newlyDestroyed, collision.ids[0], collision.ids[1])
		}

		for _, id := range newlyDestroyed {
			mask[id] = true
		}
	}

	cnt := 0
	for _, destroyed := range mask {
		if !destroyed {
			cnt++
		}
	}

	return cnt
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
		particles = append(particles, parseParticle(scanner.Text()))
	}

	minAccParticle := slices.MinFunc(particles, func(a, b Particle) int {
		return a.manhattanAcceleration() - b.manhattanAcceleration()
	})
	println(slices.Index(particles, minAccParticle))

	println(resolveCollisions(particles))
}
