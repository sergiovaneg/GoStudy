package p0238_test

import (
	"reflect"
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p0238"
)

func TestProductExceptSelf(t *testing.T) {
	if !reflect.DeepEqual(
		p0238.ProductExceptSelf([]int{1, 2, 3, 4}),
		[]int{24, 12, 8, 6}) {
		t.Fatal()
	}
	if !reflect.DeepEqual(
		p0238.ProductExceptSelf([]int{-1, 1, 0, -3, 3}),
		[]int{0, 0, 9, 0, 0}) {
		t.Fatal()
	}
}
