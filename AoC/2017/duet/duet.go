package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/sergiovaneg/GoStudy/utils"
)

const nThreads = 2
const bufferSize = 10000000

type Registry map[string]int

type Scheduler struct {
	locker [nThreads]*sync.Mutex
	cnt    [nThreads]int
	queue  [nThreads]chan int
}

func initScheduler() Scheduler {
	var s Scheduler

	for i := range nThreads {
		s.queue[i] = make(chan int, bufferSize)
		s.locker[i] = new(sync.Mutex)
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

func (s *Scheduler) run(instructions []string, src, dst int) {
	r := make(Registry)
	ptr, n := 0, len(instructions)

	for ptr >= 0 && ptr < n {
		inst := strings.Split(instructions[ptr], " ")

		switch strings.ToLower(inst[0]) {
		case "rcv":
			if len(s.queue[src]) == 0 {
				isLocked := !s.locker[dst].TryLock()
				if isLocked {
					ptr = n
					break
				}
				s.locker[dst].Unlock()
			}

			s.locker[src].Lock()
			aux, ok := <-s.queue[src]
			s.locker[src].Unlock()

			if !ok {
				ptr = n
			} else {
				r[inst[1]] = aux
			}

		case "snd":
			s.queue[dst] <- r.query(inst[1])
			s.cnt[src]++
		case "jgz":
			if r.query(inst[1]) > 0 {
				ptr += r.query(inst[2]) - 1
			}
		default:
			r.update(inst)
		}

		ptr++
	}

	close(s.queue[dst])
	s.locker[src].Lock()
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

	s := initScheduler()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		s.run(instructions, 0, 1)
		wg.Done()
	}()

	go func() {
		s.run(instructions, 1, 0)
		wg.Done()
	}()

	wg.Wait()
	println(s.cnt[1])
}
