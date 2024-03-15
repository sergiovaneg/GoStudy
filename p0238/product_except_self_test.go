package p0238_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sergiovaneg/GO_leetcode/p0238"
)

func TestProductExceptSelf(t *testing.T) {
	if !cmp.Equal(
		p0238.ProductExceptSelf([]int{1, 2, 3, 4}),
		[]int{24, 12, 8, 6}) {
		t.Fatal()
	}
	if !cmp.Equal(
		p0238.ProductExceptSelf([]int{-1, 1, 0, -3, 3}),
		[]int{0, 0, 9, 0, 0}) {
		t.Fatal()
	}
}
