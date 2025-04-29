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

func (bm *BitmaskB) updateMask(line string) {
	bm.setMask = 0
	bm.floatingMasks = make([]int, 0)

	for i, c := range line[7:] {
		if c == '0' {
			continue
		}

		mask := 0x01 << (35 - i)
		if c == '1' {
			bm.setMask |= mask
		} else {
			bm.floatingMasks = append(bm.floatingMasks, mask)
		}
	}
}

func (bm BitmaskA) updateMem(line string, m *Memory) {
	nums := regexp.MustCompile(`(\d+)`).FindAllString(line, 2)

	addr, _ := strconv.Atoi(nums[0])
	val, _ := strconv.Atoi(nums[1])

	for _, op := range bm {
		if op.dir {
			val |= op.mask
		} else {
			val &= ^op.mask
		}
	}

	(*m)[addr] = val
}

func recursiveGen(addr int, masks []int) []int {
	if len(masks) == 0 {
		return []int{addr}
	}

	return append(
		recursiveGen(addr&^masks[0], masks[1:]),
		recursiveGen(addr|masks[0], masks[1:])...)
}

func (bm BitmaskB) updateMem(line string, m *Memory) {
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
		line := scanner.Text()
		if strings.HasPrefix(line, "mask") {
			maskA.updateMask(line)
			maskB.updateMask(line)
		} else {
			maskA.updateMem(line, &memA)
			maskB.updateMem(line, &memB)
		}
	}

	println(memA.getSum())
	println(memB.getSum())
}
