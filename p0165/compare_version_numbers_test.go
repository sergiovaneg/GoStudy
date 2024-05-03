package p0165_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p0165"
)

func TestCompareVersion(t *testing.T) {
	var res int

	res = p0165.CompareVersion("1.01", "1.001")
	if res != 0 {
		t.Fatalf("Expected 0; got %v", res)
	}

	res = p0165.CompareVersion("1.0", "1.0.0")
	if res != 0 {
		t.Fatalf("Expected 0; got %v", res)
	}

	res = p0165.CompareVersion("1.0.1", "1")
	if res != 1 {
		t.Fatalf("Expected 1; got %v", res)
	}
}
