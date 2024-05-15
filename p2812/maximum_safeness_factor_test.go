package p2812_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p2812"
)

func TestMaximumSafenessFactor(t *testing.T) {
	var res int
	res = p2812.MaximumSafenessFactor([][]int{
		{1, 0, 0},
		{0, 0, 0},
		{0, 0, 1},
	})
	if res != 0 {
		t.Fatalf("Expected 0; got %v", res)
	}

	res = p2812.MaximumSafenessFactor([][]int{
		{0, 0, 1},
		{0, 0, 1},
		{0, 0, 0},
	})
	if res != 1 {
		t.Fatalf("Expected 1; got %v", res)
	}

	res = p2812.MaximumSafenessFactor([][]int{
		{0, 0, 0, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{1, 0, 0, 0},
	})
	if res != 2 {
		t.Fatalf("Expected 2; got %v", res)
	}
}
