package p1656

import "slices"

type OrderedStream struct {
	n      int
	ptr    int
	buffer []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{n: n, ptr: 0, buffer: make([]string, n)}
}

func (stream *OrderedStream) Insert(idKey int, value string) []string {
	stream.buffer[idKey-1] = value
	idx := slices.IndexFunc(stream.buffer[stream.ptr:],
		func(s string) bool { return s == "" })

	var res []string
	if idx == -1 {
		res = stream.buffer[stream.ptr:]
		stream.ptr = stream.n
	} else {
		res = stream.buffer[stream.ptr : stream.ptr+idx]
		stream.ptr += idx
	}
	return res
}
