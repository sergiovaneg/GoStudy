package utils

import (
	"bytes"
	"cmp"
	"io"
	"log"
	"os"
	"slices"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func LineCounter(r *os.File) (int, error) {
	offset, err := r.Seek(0, io.SeekCurrent)
	if err != nil {
		return 0, err
	}

	buf := make([]byte, 512*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			r.Seek(offset, io.SeekStart)
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func GCD[I Integer](a, b I) I {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM[I Integer](a, b I, integers ...I) I {
	result := a * b / GCD(a, b)

	for i := range integers {
		result = LCM(result, integers[i])
	}

	return result
}

func ExtGCD[I Integer](a, b I) (I, I, I) {
	rOld, r := a, b
	sOld, s := I(1), I(0)
	tOld, t := I(0), I(1)
	for r != 0 {
		quotient, remainder := rOld/r, rOld%r
		rOld, r = r, remainder
		sOld, s = s, sOld-quotient*s
		tOld, t = t, tOld-quotient*t
	}

	return rOld, sOld, tOld
}

func ExtLCM[I Integer](mu []I, lambda []I) (I, I) {
	muC, lambdaC := mu[0], lambda[0]

	for idx := range mu[1:] {
		g, s, _ := ExtGCD(lambdaC, lambda[idx+1])
		phaseDiff := muC - mu[idx+1]
		pdMult, pdRem := phaseDiff/g, phaseDiff%g

		if pdRem != 0 {
			log.Fatal("No synchronization possible")
		}

		lambdaNew := (lambdaC / g) * lambda[idx+1]
		muNew := (muC - s*pdMult*lambdaC) % lambdaNew
		lambdaC, muC = lambdaNew, muNew
	}

	return muC, lambdaC
}

func IPow[I Integer](a, b I) I {
	r := I(1)

	for b != 0 {
		if b&0x01 == 1 {
			r *= a
		}
		b >>= 1
		a *= a
	}

	return r
}

func SliceDifference[T comparable](a, b []T) []T {
	c := make([]T, 0, len(a))
	for _, elem := range a {
		if !slices.Contains(b, elem) {
			c = append(c, elem)
		}
	}

	return c
}

func SortedUniqueInsert[T cmp.Ordered](x []T, e T) ([]T, bool) {
	i, ok := slices.BinarySearch(x, e)

	if !ok {
		x = slices.Insert(x, i, e)
	}

	return x, ok
}

func SortedMerge[T cmp.Ordered](a, b []T) []T {
	la, lb := len(a), len(b)
	c := make([]T, 0, la+lb)
	ia, ib := 0, 0

	for ia < la && ib < lb {
		if a[ia] < b[ib] {
			c = append(c, a[ia])
			ia++
		} else {
			c = append(c, b[ib])
			ib++
		}
	}

	c = append(c, a[ia:]...)
	c = append(c, b[ib:]...)

	return c
}

func AbsInt[T Integer](x T) T {
	if x < 0 {
		return -x
	}

	return x
}

func ISqrt[T Integer](x T) T {
	if x <= 1 {
		return x
	}

	x0 := x >> 1
	x1 := (x0 + x/x0) >> 1

	for x1 < x0 {
		x0 = x1
		x1 = (x0 + x/x0) >> 1
	}

	return x0
}

func Permute[T any, S []T](s S) []S {
	if len(s) == 1 {
		return []S{s}
	}

	perms := make([]S, 0)
	for idx, elem := range s {
		aux := make(S, len(s))
		copy(aux, s)
		aux = slices.Delete(aux, idx, idx+1)
		for _, perm := range Permute(aux) {
			perms = append(perms, append([]T{elem}, perm...))
		}
	}

	return perms
}

func Choices[T any](collection [][]T) [][]T {
	if len(collection[0]) == 0 {
		return [][]T{}
	}

	if len(collection) == 1 {
		split := make([][]T, len(collection[0]))
		for idx, elem := range collection[0] {
			split[idx] = []T{elem}
		}
		return split
	}

	next := Choices(collection[1:])
	drafts := make([][]T, 0, len(collection[0])*len(next))

	for _, tail := range next {
		for _, head := range collection[0] {
			drafts = append(drafts, append([]T{head}, tail...))
		}
	}

	return drafts
}
