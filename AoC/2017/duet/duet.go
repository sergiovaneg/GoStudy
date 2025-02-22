package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const nThreads = 2

type Registry map[string]int

type Scheduler struct {
	cnt    [nThreads]int
	ptrs   [nThreads]int
	regs   [nThreads]Registry
	queues [nThreads][]int
}

func initScheduler() Scheduler {
	var s Scheduler

	for i := range nThreads {
		s.regs[i] = make(Registry)
		s.regs[i]["p"] = i
		s.queues[i] = make([]int, 0)
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

func concurrentExecution(instructions []string) int {
	s := initScheduler()
	n := len(instructions)

	for {
		isDeadlocked := true

		for src := range nThreads {
			if s.ptrs[src] < 0 || s.ptrs[src] >= n {
				continue
			}

			inst := strings.Split(instructions[s.ptrs[src]], " ")

			switch strings.ToLower(inst[0]) {
			case "rcv":
				if len(s.queues[src]) == 0 {
					continue
				}
				s.regs[src][inst[1]] = s.queues[src][0]
				s.queues[src] = s.queues[src][1:]
			case "snd":
				s.queues[(src+1)%nThreads] = append(
					s.queues[(src+1)%nThreads],
					s.regs[src].query(inst[1]),
				)
				s.cnt[src]++
			case "jgz":
				if s.regs[src].query(inst[1]) > 0 {
					s.ptrs[src] += s.regs[src].query(inst[2]) - 1
				}
			default:
				s.regs[src].update(inst)
			}

			s.ptrs[src]++
			isDeadlocked = false
		}

		if isDeadlocked {
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

	instructions := make([]string, 0)

	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	println(singleExecution(instructions))
	println(concurrentExecution(instructions))
}
