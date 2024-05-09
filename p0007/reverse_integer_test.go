package p0007_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p0007"
)

func TestReverse(t *testing.T) {
	var res int

	res = p0007.Reverse(123)
	if res != 321 {
		t.Fatalf("Expected 321; got %v", res)
	}

	res = p0007.Reverse(-123)
	if res != -321 {
		t.Fatalf("Expected -321; got %v", res)
	}

	res = p0007.Reverse(120)
	if res != 21 {
		t.Fatalf("Expected 21; got %v", res)
	}
}
