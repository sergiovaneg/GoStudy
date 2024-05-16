package p0006_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/LeetCode/p0006"
)

func TestConvert(t *testing.T) {
	var res string

	res = p0006.Convert("PAYPALISHIRING", 3)
	if res != "PAHNAPLSIIGYIR" {
		t.Fatalf("Expected 'PAHNAPLSIIGYIR'; got %v", res)
	}

	res = p0006.Convert("PAYPALISHIRING", 4)
	if res != "PINALSIGYAHRPI" {
		t.Fatalf("Expected 'PINALSIGYAHRPI'; got %v", res)
	}

	res = p0006.Convert("A", 1)
	if res != "A" {
		t.Fatalf("Expected 'A'; got %v", res)
	}
}
