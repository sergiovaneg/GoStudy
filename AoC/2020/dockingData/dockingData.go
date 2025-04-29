package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Memory map[int]int
type OperationA struct {
	mask int
	dir  bool
}
type BitmaskA []OperationA

type BitmaskB struct {
	setMask       int
	floatingMasks []int
}

func (bm *BitmaskA) updateMask(line string) {
	*bm = make(BitmaskA, 0)
	for i, c := range line[7:] {
		if c == 'X' {
			continue
		}
		*bm = append(*bm, OperationA{0x01 << (35 - i), c == '1'})
	}
}

func recursiveGen(addr int, masks []int) []int {
	if len(masks) == 0 {
		return []int{addr}
	}

	addrs := make([]int, 0, 0x01<<len(masks))

	addrs = append(addrs, recursiveGen(addr&^masks[0], masks[1:])...)
	addrs = append(addrs, recursiveGen(addr|masks[0], masks[1:])...)

	return addrs
}

func (bm *BitmaskB) updateMask(line string) {
	bm.setMask = 0
	bm.floatingMasks = make([]int, 0)

	for i, c := range line[7:] {
		if c == '0' {
			continue
		}
		if c == '1' {
			bm.setMask |= 0x01 << (35 - i)
			continue
		}

		bm.floatingMasks = append(bm.floatingMasks, 0x01<<(35-i))
	}
}

func (m *Memory) updateA(line string, bm BitmaskA) {
	nums := regexp.MustCompile(`(\d+)`).FindAllString(line, 2)
	val, _ := strconv.Atoi(nums[1])

	for _, op := range bm {
		if op.dir {
			val |= op.mask
		} else {
			val &= ^op.mask
		}
	}

	pos, _ := strconv.Atoi(nums[0])
	(*m)[pos] = val
}

func (m *Memory) updateB(line string, bm BitmaskB) {
	nums := regexp.MustCompile(`(\d+)`).FindAllString(line, 2)

	src, _ := strconv.Atoi(nums[0])
	val, _ := strconv.Atoi(nums[1])

	for _, addr := range recursiveGen(src|bm.setMask, bm.floatingMasks) {
		(*m)[addr] = val
	}
}

func (m Memory) getSum() int {
	res := 0
	for _, val := range m {
		res += val
	}
	return res
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	maskA := make(BitmaskA, 0)
	maskB := BitmaskB{
		floatingMasks: make([]int, 0),
	}
	memA, memB := make(Memory), make(Memory)

	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "mask") {
			maskA.updateMask(scanner.Text())
			maskB.updateMask(scanner.Text())
		} else {
			memA.updateA(scanner.Text(), maskA)
			memB.updateB(scanner.Text(), maskB)
		}
	}

	println(memA.getSum())
	println(memB.getSum())
}
