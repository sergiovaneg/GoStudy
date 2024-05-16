package p3075_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/LeetCode/p3075"
)

func TestMaximumHappinessSum(t *testing.T) {
	var res int64

	res = p3075.MaximumHappinessSum([]int{1, 2, 3}, 2)
	if res != 4 {
		t.Fatalf("Expected 4; got %v", res)
	}

	res = p3075.MaximumHappinessSum([]int{1, 1, 1, 1}, 2)
	if res != 1 {
		t.Fatalf("Expected 1; got %v", res)
	}

	res = p3075.MaximumHappinessSum([]int{2, 3, 4, 5}, 1)
	if res != 5 {
		t.Fatalf("Expected 5; got %v", res)
	}
}
