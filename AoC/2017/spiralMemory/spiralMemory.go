package main

import "math"

const target int = 289326

type Memory map[[2]int]int

func getManhattanDistance(addr int) int {
	n := int(math.Ceil((math.Sqrt(float64(addr)) - 1.) / 2.))
	ref := (n<<1 + 1)
	ref *= ref

	for ref > addr {
		ref -= n << 1
	}

	ref += n
	if target > ref {
		return (addr - ref) + n
	} else {
		return (ref - addr) + n
	}
}

func (m *Memory) updateAt(x [2]int) {
	var acc int
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			acc += (*m)[[2]int{x[0] + dx, x[1] + dy}]
		}
	}
	(*m)[x] = acc
}

func firstLargest(thr int) int {
	mem := Memory{[2]int{0, 0}: 1}
	n, x := 0, [2]int{0, 0}
	for {
		n++ // Next ring

		x[1]++ //Move right
		mem.updateAt(x)
		if mem[x] > thr {
			return mem[x]
		}

		for range n<<1 - 1 {
			x[0]-- // Move up
			mem.updateAt(x)

			if mem[x] > thr {
				return mem[x]
			}
		}

		// Other directions
		for _, d := range [3][2]int{{0, -1}, {1, 0}, {0, 1}} {
			for range n << 1 {
				x[0] += d[0]
				x[1] += d[1]
				mem.updateAt(x)

				if mem[x] > thr {
					return mem[x]
				}
			}
		}
	}
}

func main() {
	println(getManhattanDistance(target))
	println(firstLargest(target))
}
