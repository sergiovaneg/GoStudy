package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/sergiovaneg/GoStudy/utils"
)

const nThreads = 2

type Registry map[string]int

type Scheduler struct {
	cnt   [nThreads]int
	flags [nThreads]bool
	ptrs  [nThreads]int
	reg   [nThreads]Registry
	queue [nThreads][]int
}

func initScheduler() Scheduler {
	var s Scheduler

	for i := range nThreads {
		s.reg[i] = make(Registry)
		s.queue[i] = make([]int, 0)
	}

	return s
}

func (r Registry) query(key string) int {
	aux, err := strconv.Atoi(key)
	if err != nil {
		aux = r[key]
	}

	return aux
}

func (r *Registry) update(inst []string) {
	aux := r.query(inst[2])

	switch strings.ToLower(inst[0]) {
	case "set":
		(*r)[inst[1]] = aux
	case "add":
		(*r)[inst[1]] += aux
	case "mul":
		(*r)[inst[1]] *= aux
	case "mod":
		(*r)[inst[1]] %= aux
	}
}

func singleExecution(instructions []string) int {
	r := make(Registry)
	ptr, n := 0, len(instructions)

	for ptr >= 0 && ptr < n {
		inst := strings.Split(instructions[ptr], " ")

		switch strings.ToLower(inst[0]) {
		case "rcv":
			if r.query(inst[1]) > 0 {
				ptr = n
			}
		case "snd":
			r[""] = r.query(inst[1])
		case "jgz":
			if r.query(inst[1]) > 0 {
				ptr += r.query(inst[2]) - 1
			}
		default:
			r.update(inst)
		}

		ptr++
	}

	return r[""]
}

func (s Scheduler) isDeadlocked() bool {
	for _, f := range s.flags {
		if f {
			return false
		}
	}

	return true
}

func concurrentExecution(instructions []string) int {
	s := initScheduler()
	n := len(instructions)

	for {
		s.flags = [nThreads]bool{}

		for src := range nThreads {
			if s.ptrs[src] < 0 || s.ptrs[src] >= n {
				continue
			}

			inst := strings.Split(instructions[s.ptrs[src]], " ")

			switch strings.ToLower(inst[0]) {
			case "rcv":
				if len(s.queue[src]) == 0 {
					continue
				}
				s.reg[src][inst[1]] = s.queue[src][0]
				s.queue[src] = s.queue[src][1:]
			case "snd":
				s.queue[(src+1)%nThreads] = append(
					s.queue[(src+1)%nThreads],
					s.reg[src].query(inst[1]),
				)
				s.cnt[src]++
			case "jgz":
				if s.reg[src].query(inst[1]) > 0 {
					s.ptrs[src] += s.reg[src].query(inst[2]) - 1
				}
			default:
				s.reg[src].update(inst)
			}

			s.ptrs[src]++
			s.flags[src] = true
		}

		if s.isDeadlocked() {
			break
		}
	}

	return s.cnt[nThreads-1]
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n, _ := utils.LineCounter(file)

	instructions := make([]string, n)

	for idx := 0; scanner.Scan(); idx++ {
		instructions[idx] = scanner.Text()
	}

	println(singleExecution(instructions))
	println(concurrentExecution(instructions))
}
