package p1915_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/LeetCode/p1915"
)

func TestWonderfulSubstrings(t *testing.T) {
	var res int64

	res = p1915.WonderfulSubstrings("aba")
	if res != 4 {
		t.Fatalf("Expected 4; got %v", res)
	}

	res = p1915.WonderfulSubstrings("aabb")
	if res != 9 {
		t.Fatalf("Expected 9; got %v", res)
	}

	res = p1915.WonderfulSubstrings("he")
	if res != 2 {
		t.Fatalf("Expected 2; got %v", res)
	}

	res = p1915.WonderfulSubstrings("jfdghjjejjbghchijfj")
	if res != 29 {
		t.Fatalf("Expected 29; got %v", res)
	}
}
