package p1544_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/LeetCode/p1544"
)

func TestMakeGood(t *testing.T) {
	var res string

	res = p1544.MakeGood("leEeetcode")
	if res != "leetcode" {
		t.Fatalf("Expected 'leetcode', got '%v'", res)
	}

	res = p1544.MakeGood("abBAcC")
	if res != "" {
		t.Fatalf("Expected '', got '%v'", res)
	}

	res = p1544.MakeGood("s")
	if res != "s" {
		t.Fatalf("Expected 's', got '%v'", res)
	}
}
