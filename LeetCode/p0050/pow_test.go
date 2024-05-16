package p0050_test

import (
	"math"
	"testing"

	"github.com/sergiovaneg/GO_leetcode/LeetCode/p0050"
)

const eps float64 = 1e-6

func TestPow(t *testing.T) {
	var res float64

	res = p0050.MyPow(2., 10)
	if math.Abs(res-1024.) > eps {
		t.Fatalf("Expected %v; got %v", 1024., res)
	}

	res = p0050.MyPow(2.1, 3)
	if math.Abs(res-9.26100) > eps {
		t.Fatalf("Expected %v; got %v", 9.26100, res)
	}

	res = p0050.MyPow(2., -2)
	if math.Abs(res-0.25) > eps {
		t.Fatalf("Expected %v; got %v", 0.25, res)
	}
}
