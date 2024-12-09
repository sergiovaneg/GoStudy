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

func compactHash(fragmentedDisk string) int {
	l := len(fragmentedDisk)
	acc, idx := 0, 0
	i, j := 0, l-2+l&0x01
	head_id, tail_id := 0, (l-1)>>1

	queue := fragmentedDisk[j] - '0'
	for ; i < j; i++ {
		aux := fragmentedDisk[i] - '0'
		if i&0x01 == 0 {
			for range aux {
				acc += idx * head_id
				idx++
			}
			head_id++
		} else {
			for range aux {
				if queue == 0 {
					j -= 2
					tail_id--
					queue = fragmentedDisk[j] - '0'
				}
				acc += idx * tail_id
				idx++
				queue--
			}
		}
	}

	for range queue {
		acc += idx * tail_id
		idx++
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
			}); idx != -1 && freeBlocks[idx][0] > file[0] {
			file[0] = freeBlocks[idx][0]
			if freeBlocks[idx][1] == file[1] {
				freeBlocks = slices.Delete(freeBlocks, idx, idx+1)
			} else {
				freeBlocks[idx][0] += file[1]
				freeBlocks[idx][1] -= file[1]
			}
		}

		acc += (id * file[1] * (file[0]<<1 + file[1] - 1)) >> 1
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
