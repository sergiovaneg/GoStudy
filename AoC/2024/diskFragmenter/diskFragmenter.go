package main

import (
	"bufio"
	"log"
	"os"
	"slices"
)

func getBlocks(fragmentedDisk string) ([][2]int, [][2]int) {
	l, pos := len(fragmentedDisk), 0
	res := [2][][2]int{make([][2]int, l>>1+l&0x01), make([][2]int, l>>1)}

	for idx, c := range fragmentedDisk {
		aux := int(c - '0')
		res[idx&0x01][idx>>1] = [2]int{pos, aux}
		pos += aux
	}

	return res[0], res[1]
}

func hashContiguous(id, pos, cap int) int {
	return (id * cap * (pos<<1 + cap - 1)) >> 1
}

func compactHash(fragmentedDisk string) int {
	files, freeBlocks := getBlocks(fragmentedDisk)
	slices.Reverse(files)

	acc, id := 0, len(files)-1
	for files[0][0] > freeBlocks[0][0] {
		cap, pos := min(files[0][1], freeBlocks[0][1]), freeBlocks[0][0]
		acc += hashContiguous(id, pos, cap)

		if files[0][1] == cap {
			files = files[1:]
			id--
		} else {
			files[0][1] -= cap
		}

		if freeBlocks[0][1] == cap {
			freeBlocks = freeBlocks[1:]
		} else {
			freeBlocks[0][0] += cap
			freeBlocks[0][1] -= cap
		}
	}

	for _, file := range files {
		acc += hashContiguous(id, file[0], file[1])
		id--
	}

	return acc
}

func blockCompactHash(fragmentedDisk string) int {
	files, freeBlocks := getBlocks(fragmentedDisk)

	acc := 0
	for id := len(files) - 1; id >= 0; id-- {
		file := files[id]

		if idx := slices.IndexFunc(
			freeBlocks,
			func(free [2]int) bool {
				return free[1] >= file[1]
			}); idx != -1 && freeBlocks[idx][0] < file[0] {
			file[0] = freeBlocks[idx][0]

			if freeBlocks[idx][1] == file[1] {
				freeBlocks = slices.Delete(freeBlocks, idx, idx+1)
			} else {
				freeBlocks[idx][0] += file[1]
				freeBlocks[idx][1] -= file[1]
			}
		}

		acc += hashContiguous(id, file[0], file[1])
	}

	return acc
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	fragDisk := scanner.Text()

	println(compactHash(fragDisk))
	println(blockCompactHash(fragDisk))
}
