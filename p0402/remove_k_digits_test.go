package p0402_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p0402"
)

func TestRemoveKDigits(t *testing.T) {
	var res string

	res = p0402.RemoveKdigits("1432219", 3)
	if res != "1219" {
		t.Fatalf("Expected '1219'; got '%v'", res)
	}

	res = p0402.RemoveKdigits("10200", 1)
	if res != "200" {
		t.Fatalf("Expected '200'; got '%v'", res)
	}

	res = p0402.RemoveKdigits("10", 2)
	if res != "0" {
		t.Fatalf("Expected '0'; got '%v'", res)
	}
}
