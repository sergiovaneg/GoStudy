package p2073_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p2073"
)

func TestTimeRequiredToBuy(t *testing.T) {
	var res int

	res = p2073.TimeRequiredToBuy([]int{2, 3, 2}, 2)
	if res != 6 {
		t.Fatalf("Expected 6; got %v", res)
	}

	res = p2073.TimeRequiredToBuy([]int{5, 1, 1, 1}, 0)
	if res != 8 {
		t.Fatalf("Expected 8; got %v", res)
	}
}
