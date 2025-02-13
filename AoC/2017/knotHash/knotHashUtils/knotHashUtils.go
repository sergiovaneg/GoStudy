package knothashutils

import "fmt"

const listSize = 256

type State struct {
	Lst [listSize]int
	pos int
	skp int
}

func InitState() State {
	var s State
	for i := range listSize {
		s.Lst[i] = i
	}

	return s
}

func (s *State) iter(l int) {
	aux, j := make([]int, l), s.pos

	for i := range l {
		aux[i] = s.Lst[j]

		if j == listSize-1 {
			j = 0
		} else {
			j++
		}
	}

	for i := range l {
		if j == 0 {
			j = listSize - 1
		} else {
			j--
		}

		s.Lst[j] = aux[i]
	}

	s.pos = (s.pos + l + s.skp) % listSize
	s.skp++
}

func (s *State) SparseHash(iLengths []int, nRounds int) {
	for range nRounds {
		for _, l := range iLengths {
			s.iter(l)
		}
	}
}

func (s State) DenseHash() string {
	hash := ""

	for i := 0; i < listSize; i += 16 {
		var aux int
		for j := range 16 {
			aux ^= s.Lst[i+j]
		}
		hash += fmt.Sprintf("%0.2x", aux)
	}

	return hash
}

func KnotHash(text string) string {
	s := InitState()
	iLengths := make([]int, 0)

	for _, c := range text {
		iLengths = append(iLengths, int(c))
	}

	for _, v := range [5]int{17, 31, 73, 47, 23} {
		iLengths = append(iLengths, v)
	}

	s.SparseHash(iLengths, 64)

	return s.DenseHash()
}
