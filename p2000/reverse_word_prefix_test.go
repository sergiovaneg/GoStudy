package p2000_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p2000"
)

func TestReversePrefix(t *testing.T) {
	var res string

	res = p2000.ReversePrefix("abcdefd", 'd')
	if res != "dcbaefd" {
		t.Fatalf("Expected 'dcbaefd'; got '%v'", res)
	}

	res = p2000.ReversePrefix("xyxzxe", 'z')
	if res != "zxyxxe" {
		t.Fatalf("Expected 'zxyxxe'; got '%v'", res)
	}

	res = p2000.ReversePrefix("abcd", 'z')
	if res != "abcd" {
		t.Fatalf("Expected 'abcd'; got '%v'", res)
	}
}
