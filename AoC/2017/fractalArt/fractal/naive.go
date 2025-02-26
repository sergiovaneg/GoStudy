// go:build naive

package fractal

import (
	"slices"
	"strings"
)

type Fractal [][]byte
type Ruleset map[string]Fractal

func (f Fractal) serialize() string {
	rows := make([]string, len(f))

	for i, row := range f {
		rows[i] = string(row)
	}

	return strings.Join(rows, "/")
}

func deserialize(serial string) Fractal {
	f := make(Fractal, 0)

	for _, row := range strings.Split(serial, "/") {
		f = append(f, []byte(row))
	}

	return f
}

func (r *Ruleset) Update(line string) {
	kv := strings.Split(line, " => ")
	(*r)[kv[0]] = deserialize(kv[1])
}

func makeEmpty(n int) Fractal {
	f := make(Fractal, n)
	for i := range n {
		f[i] = make([]byte, n)
	}
	return f
}

func (f Fractal) mirror() Fractal {
	fNew := make(Fractal, len(f))

	for i, row := range f {
		fNew[i] = slices.Clone(row)
		slices.Reverse(fNew[i])
	}

	return fNew
}

func (f Fractal) rotate() Fractal {
	n := len(f)
	fNew := makeEmpty(n)

	for i := range n {
		for j := range n {
			fNew[i][j] = f[j][n-i-1]
		}
	}

	return fNew
}

func (r Ruleset) transform(f Fractal) Fractal {
	fm := f.mirror()

	for range 4 {
		if aux, ok := r[f.serialize()]; ok {
			return aux
		}

		if aux, ok := r[fm.serialize()]; ok {
			return aux
		}

		f = f.rotate()
		fm = fm.rotate()
	}

	return nil
}
